package controller

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/anacrolix/torrent"
	"github.com/asdine/storm"
	"github.com/gnur/golpje/database"
	"github.com/gnur/golpje/downloader"
	"github.com/gnur/golpje/events"
	pb "github.com/gnur/golpje/golpje"
	"github.com/gnur/golpje/searcher"
	"github.com/gnur/golpje/shows"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

// controller is a stub
type controller struct {
	Searchresults   chan searcher.Searchresult
	DownloadChannel chan downloader.Download
	config          *viper.Viper
}

// Start commences the controller
func Start(config *viper.Viper) error {
	var con controller
	con.config = config

	database.Conn, _ = storm.Open(con.config.GetString("database_file"))
	defer database.Conn.Close()
	lis, err := net.Listen("tcp", con.config.GetString("port"))

	if err != nil {
		fmt.Println("IM NOTLISTENING!")
		fmt.Println(err.Error())
		return nil
	}
	con.Searchresults = make(chan searcher.Searchresult)
	con.DownloadChannel = make(chan downloader.Download, 40) //buffered channel so it doesn't block and queues new downloads
	go downloader.Start(con.DownloadChannel)
	if con.config.GetBool("search_enabled") {
		go searcher.Start(con.Searchresults, con.config.GetDuration("search_interval"))
	}
	go con.resulthandler()
	s := grpc.NewServer()
	pb.RegisterGolpjeServer(s, &con)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve")
	}

	return nil
}

func (con *controller) resulthandler() {
	for res := range con.Searchresults {
		fmt.Println("--------------")
		if res.Seeders < 10 || !strings.Contains(res.Title, "264") {
			fmt.Println("too little seeders or not 264")
			continue
		}
		fmt.Println(res.ShowID, res.Title)
		show, err := shows.GetFromID(res.ShowID)
		if err != nil {
			fmt.Println("continuing ", err.Error())
			continue
		}
		shouldDownload, err := show.ShouldDownload(res.Title)

		if !shouldDownload {
			fmt.Println("not downloading..")
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("yes: ")
		fmt.Println(res.Title)
		downloadID, err := show.AddDownload(res.Title, res.Magnetlink)
		if err != nil {
			fmt.Println("got an error..")
			fmt.Println(err.Error())
			continue
		}

		events.New(fmt.Sprintf("Starting download of %s", res.Title), []string{res.ShowID, downloadID})
		fmt.Println("starting download")
		fmt.Println(downloadID)
		resultChannel := make(chan downloader.Result)
		downloadPath := fmt.Sprintf("%s/%s", con.config.GetString("download_path"), downloadID)
		dl := downloader.Download{
			Magnetlink:    res.Magnetlink,
			DownloadDir:   downloadPath,
			ResultChannel: resultChannel,
		}
		fmt.Println("sending download to channel")
		con.DownloadChannel <- dl
		fmt.Println("waiting for result")
		downloadResult := <-resultChannel
		if !downloadResult.Completed {
			fmt.Println("Download did not complete")
			fmt.Println(downloadResult.Error)
			events.New(fmt.Sprintf("Download of %s failed; %s", res.Title, downloadResult.Error), []string{res.ShowID, downloadID})
			show.SetDownloadFailed(res.Title)
			continue
		}

		fmt.Println("download completed")
		var largestFile torrent.File
		var largest int64
		largest = 0
		for _, f := range downloadResult.Files {
			fmt.Println(f.Path())
			if f.Length() > largest {
				largest = f.Length()
				largestFile = f
			}
		}
		fmt.Println("setting as downloaded: ", res.Title)
		showPath := show.Path(con.config.GetString("shows_path"))
		targetDir := show.GetSeasonDir(res.Title, showPath)
		targetName := filepath.Join(targetDir, filepath.Base(largestFile.Path()))
		sourceName := filepath.Join(downloadPath, largestFile.Path())
		err = os.MkdirAll(filepath.Dir(targetName), 0777)
		if err != nil {
			fmt.Println("mkdirall error: ", err.Error())
			continue
		}

		err = os.Rename(sourceName, targetName)
		if err != nil {
			fmt.Println("rename error: ", err.Error())
			continue
		}
		events.New(fmt.Sprintf("Completed download of %s", res.Title), []string{res.ShowID, downloadID})
		show.SetDownloaded(res.Title)
	}
}

func (con *controller) GetEvents(ctx context.Context, in *pb.EventRequest) (*pb.ProtoEvents, error) {
	var ev []events.Event
	var err error
	if in.All {
		ev, err = events.All()
	} else {
		ev, err = events.After(time.Unix(0, in.Since))
	}
	if err != nil {
		fmt.Println(err.Error())
		return &pb.ProtoEvents{}, err
	}

	var retEvents []*pb.ProtoEvent
	for _, event := range ev {
		retEvents = append(retEvents, event.ToProto())
	}
	return &pb.ProtoEvents{
		Events: retEvents,
	}, nil
}

func (con *controller) GetShows(ctx context.Context, in *pb.ShowRequest) (*pb.ProtoShows, error) {
	var resp pb.ProtoShows
	allShows, _ := shows.All()
	for _, show := range allShows {
		fmt.Println(show.ID, show.Active)
		if (!in.Onlyactive || (in.Onlyactive && show.Active)) && (in.Name == "" || in.Name == show.Name) {
			resp.Shows = append(resp.Shows, show.ToProto())
		}
	}

	return &resp, nil
}

func (con *controller) AddShow(ctx context.Context, in *pb.ProtoShow) (*pb.AddShowResponse, error) {
	var resp pb.AddShowResponse
	uuid, err := shows.New(in.Name, in.Regexp, in.Seasonal, in.Active, in.Minimal)
	if err != nil {
		resp.Error = err.Error()
	} else {
		s, _ := shows.GetFromID(uuid)
		resp.Show = s.ToProto()
	}

	return &resp, nil
}

func (con *controller) DelShow(ctx context.Context, in *pb.ProtoShow) (*pb.AddShowResponse, error) {
	var resp pb.AddShowResponse
	show, err := shows.GetFromID(in.ID)
	if err != nil {
		resp.Error = err.Error()
	} else {
		resp.Show = show.ToProto()
		show.Delete()
	}

	return &resp, nil
}

func (con *controller) GetEpisodes(ctx context.Context, in *pb.EpisodeRequest) (*pb.ProtoEpisodes, error) {
	return &pb.ProtoEpisodes{}, nil
}

func (con *controller) SyncShow(ctx context.Context, in *pb.SyncShowRequest) (*pb.SyncShowResponse, error) {
	//get show
	//remove all episodes
	//loop over all files and add episodes as downloaded
	show, err := shows.GetFromID(in.ShowID)
	if err != nil {
		return &pb.SyncShowResponse{
			Error: err.Error(),
		}, nil
	}
	show.DeleteAllEpisodes()
	showdir := show.Path(con.config.GetString("shows_path"))
	var totalSynced int64
	totalSynced = 0
	filepath.Walk(showdir, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			if !info.IsDir() {
				show.AddEpisode(info.Name())
				totalSynced++
			}
		}
		return nil
	})

	return &pb.SyncShowResponse{
		Success:       true,
		FoundEpisodes: totalSynced,
		Error:         "",
	}, nil
}

func (con *controller) AddEpisode(ctx context.Context, in *pb.ProtoEpisode) (*pb.AddEpisodeResponse, error) {
	return &pb.AddEpisodeResponse{}, nil
}

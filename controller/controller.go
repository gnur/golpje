package controller

import (
	"fmt"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/gnur/golpje/downloader"
	"github.com/gnur/golpje/events"
	pb "github.com/gnur/golpje/golpje"
	"github.com/gnur/golpje/searcher"
	"github.com/gnur/golpje/shows"
	"golang.org/x/net/context"
)

const (
	port = ":3222"
)

// controller is a stub
type controller struct {
	Searchresults   chan searcher.Searchresult
	DownloadChannel chan downloader.Download
}

// Start commences the controller
func Start() error {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println("IM NOTLISTENING!")
		fmt.Println(err.Error())
		return nil
	}
	var con controller
	con.Searchresults = make(chan searcher.Searchresult)
	con.DownloadChannel = make(chan downloader.Download)
	go downloader.Start(con.DownloadChannel)
	go searcher.Start(con.Searchresults)
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
		if res.Seeders > 10 && strings.Contains(res.Title, "264") {
			fmt.Println(res.ShowID, res.Title)
			show, err := shows.GetFromID(res.ShowID)
			if err != nil {
				fmt.Println("continuing")
				continue
			}
			if show.ShouldDownload(res.Title) {
				fmt.Println("yes: ")
				fmt.Println(res.Title)
				downloadID, err := show.AddDownload(res.Title, res.Magnetlink)
				if err == nil {
					events.New(fmt.Sprintf("Starting download of%s", res.Title), []string{res.ShowID, downloadID})
					fmt.Println("starting download")
					fmt.Println(downloadID)
				}
			}
		}
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

func (con *controller) AddEpisode(ctx context.Context, in *pb.ProtoEpisode) (*pb.AddEpisodeResponse, error) {
	return &pb.AddEpisodeResponse{}, nil
}

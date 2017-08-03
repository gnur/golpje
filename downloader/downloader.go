package downloader

import (
	"errors"
	"fmt"
	"time"

	"github.com/anacrolix/torrent"
)

//Result holds the result of a download attempt
type Result struct {
	Files     []torrent.File
	Completed bool
	Error     error
}

//Download is the struct that holds everything needed for downloading magnetlinks
type Download struct {
	Magnetlink    string
	DownloadDir   string
	ResultChannel chan Result
}

//Start starts the downloadchannel
func Start(DownloadChannel chan Download) {
	for download := range DownloadChannel {
		result, err := downloadMagnetLink(download.Magnetlink, download.DownloadDir)
		if err != nil {
			download.ResultChannel <- Result{
				Files:     []torrent.File{},
				Completed: false,
				Error:     err,
			}
		} else {
			download.ResultChannel <- Result{
				Files:     result.Files(),
				Completed: true,
				Error:     nil,
			}
			result.Drop()
		}
	}
}

func downloadMagnetLink(magnetlink, targetDirectory string) (*torrent.Torrent, error) {
	if magnetlink[:7] != "magnet:" {
		return nil, errors.New("invalid magnetlink")
	}
	cfg := torrent.Config{
		DataDir: targetDirectory,
	}
	c, _ := torrent.NewClient(&cfg)
	defer c.Close()
	t, _ := c.AddMagnet(magnetlink)
	<-t.GotInfo()
	t.DownloadAll()
	completedChan := make(chan bool, 1)
	timeoutChan := make(chan bool, 1)
	go func() {
		c.WaitAll()
		completedChan <- true
	}()
	go func() {
		time.Sleep(20 * time.Minute)
		timeoutChan <- true
	}()
	select {
	case <-completedChan:
		fmt.Println("it worked, torrent downloaded")
		return t, nil
	case <-timeoutChan:
		fmt.Println("timeout occurred.. shit has hit the fan")
		return nil, errors.New("timeout")
	}

}

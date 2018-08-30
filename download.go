package main

import (
	"bufio"
	"bytes"
	"errors"
	"strings"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/sirupsen/logrus"
)

//Result holds the result of a download attempt
type Result struct {
	Files     []*torrent.File
	Completed bool
	Error     error
}

// Download starts a download and returns the result
func Download(magnetlink, targetDirectory string) Result {
	result, err := downloadMagnetLink(magnetlink, targetDirectory)
	if err == nil {
		defer result.Drop()
		return Result{
			Files:     result.Files(),
			Completed: true,
			Error:     nil,
		}
	}
	return Result{
		Completed: false,
		Error:     err,
	}
}

func downloadMagnetLink(magnetlink, targetDirectory string) (*torrent.Torrent, error) {
	if magnetlink[:7] != "magnet:" {
		return nil, errors.New("invalid magnetlink")
	}
	ctxLog := logrus.WithFields(logrus.Fields{
		"magnetlink":      magnetlink,
		"targetDirectory": targetDirectory,
		"scope":           "downloadMagnetLink",
	})
	cfg := torrent.ClientConfig{
		DataDir: targetDirectory,
	}
	ctxLog.Info("Downloading new torrent")
	ctxLog.Debug("starting torrent client")
	c, _ := torrent.NewClient(&cfg)
	defer c.Close()
	ctxLog.Debug("adding magnetlink to client")
	t, err := c.AddMagnet(magnetlink)
	if err != nil {
		return nil, err
	}
	<-t.GotInfo()
	ctxLog.Debug("got info from DHT")
	t.DownloadAll()
	completedChan := make(chan bool, 1)
	timeoutChan := make(chan bool, 1)
	stopStatusChan := make(chan bool, 1)
	ticker := time.NewTicker(time.Second * 30)
	ctxLog.Debug("waiting for download to complete or timeout to occur")
	go func() {
		for {
			select {
			case <-stopStatusChan:
				ticker.Stop()
				close(stopStatusChan)
				return
			case <-ticker.C:
				a := new(bytes.Buffer)
				c.WriteStatus(a)
				scanner := bufio.NewScanner(a)
				for scanner.Scan() {
					line := scanner.Text()
					if strings.Contains(line, "bytes (") {
						logrus.Debug(line)
					}
				}
			}
		}
	}()

	go func() {
		c.WaitAll()
		completedChan <- true
	}()
	go func() {
		time.Sleep(60 * time.Minute)
		timeoutChan <- true
	}()

	defer func() {
		stopStatusChan <- true
	}()

	select {
	case <-completedChan:
		return t, nil
	case <-timeoutChan:
		ctxLog.Warning("timeout while downloading")
		return nil, errors.New("timeout")
	}

}

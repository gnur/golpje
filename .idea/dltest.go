package main

import (
	"fmt"

	"github.com/gnur/golpje/downloader"
)

func main() {
	resultChannel := make(chan downloader.Result)
	downloadChannel := make(chan downloader.Download, 10)

	fmt.Println("starting downloader")

	go downloader.Start(downloadChannel)

	dl := downloader.Download{
		Magnetlink:    "magnet:?xt=urn:btih:fa0173ca923eecf1a7eb77429f6daae6ff55c836&dn=The.Simpsons.S28E18.HDTV.x264-SVA&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Fzer0day.ch%3A1337&tr=udp%3A%2F%2Fopen.demonii.com%3A1337&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Fexodus.desync.com%3A6969",
		DownloadDir:   "/tmp/dltest",
		ResultChannel: resultChannel,
	}
	fmt.Println("sending download on channel")
	downloadChannel <- dl
	fmt.Println("Waiting for result...")
	res := <-resultChannel
	fmt.Println("done?")
	fmt.Println(res.Completed)
	fmt.Println(res.Error)
	if res.Error == nil {
		for _, f := range res.Files {
			fmt.Println(f.Path())
		}
	}

}

package downloader

import (
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
	<-DownloadChannel
}

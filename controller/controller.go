package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/anacrolix/torrent"
	"github.com/asdine/storm"
	"github.com/gnur/golpje/downloader"
	"github.com/gnur/golpje/searcher"
	"github.com/gnur/golpje/shows"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// controller is a stub
type controller struct {
	Searchresults   chan searcher.Searchresult
	DownloadChannel chan downloader.Download
	config          *viper.Viper
	db              *storm.DB
	logger          *logrus.Entry
}

// Start commences the controller
func Start(config *viper.Viper) error {
	var con controller
	var err error
	con.config = config

	con.logger = logrus.WithFields(logrus.Fields{
		"scope": "controller",
	})

	con.db, err = storm.Open(con.config.GetString("database_file"))
	defer con.db.Close()
	if err != nil {
		con.logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warning("Could not open database file")
		return err
	}

	if con.config.GetBool("metrics_enabled") {
		con.logger.Info("Starting metrics server")
		go func() {
			http.Handle(con.config.GetString("metrics_path"), promhttp.Handler())
			log.Fatal(http.ListenAndServe(con.config.GetString("metrics_port"), nil))
		}()
	}

	if err != nil {
		con.logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warning("Could not start listening")
		return nil
	}
	con.Searchresults = make(chan searcher.Searchresult)
	con.DownloadChannel = make(chan downloader.Download, 40) //buffered channel so it doesn't block and queues new downloads
	if con.config.GetBool("search_enabled") {
		var sm searcher.Searchmetrics
		if con.config.GetBool("metrics_enabled") {
			sm = searcher.Searchmetrics{
				Enabled: true,
				Searches: prometheus.NewCounter(
					prometheus.CounterOpts{
						Name: "golpje_searches",
						Help: "total number of searches",
					},
				),
				FailedSearches: prometheus.NewCounter(
					prometheus.CounterOpts{
						Name: "golpje_failed_searches",
						Help: "total number of searches that failed",
					},
				),
				SearchResults: prometheus.NewCounter(
					prometheus.CounterOpts{
						Name: "golpje_search_results",
						Help: "total number of results that have been found",
					},
				),
			}
			prometheus.MustRegister(sm.Searches)
			prometheus.MustRegister(sm.FailedSearches)
			prometheus.MustRegister(sm.SearchResults)
		} else {
			sm = searcher.Searchmetrics{
				Enabled: false,
			}
		}
		con.logger.Info("Starting search goroutine")
		go searcher.Start(con.db, con.config.GetString("piratebay_url"), con.Searchresults, con.config.GetDuration("search_interval"), sm)
	}
	con.logger.Info("Starting resulthandling goroutine")
	go con.resulthandler()

	return nil
}

func (con *controller) resulthandler() {
	var failedDownloadsCounter prometheus.Counter
	var completedDownloadsCounter prometheus.Counter
	if con.config.GetBool("metrics_enabled") {
		failedDownloadsCounter = prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "golpje_failed_downloads",
				Help: "number of failed downloads",
			},
		)
		completedDownloadsCounter = prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "golpje_completed_downloads",
				Help: "number of completed downloads",
			},
		)
		prometheus.MustRegister(completedDownloadsCounter)
		prometheus.MustRegister(failedDownloadsCounter)
	}
	for res := range con.Searchresults {
		resLogger := con.logger.WithFields(logrus.Fields{
			"title":      res.Title,
			"seeders":    res.Seeders,
			"magnetlink": res.Magnetlink,
			"showid":     res.ShowID,
		})
		if res.Seeders < 10 {
			resLogger.Info("not downloading, not enough seeders")
			continue
		}
		if !validCodec(res.Title) {
			resLogger.Info("not downloading, no valid codec in title")
			continue
		}
		show, err := shows.GetFromID(con.db, res.ShowID)
		if err != nil {
			resLogger.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Info("could not extract showname")
			continue
		}
		shouldDownload, err := show.ShouldDownload(con.db, res.Title)

		if !shouldDownload {
			resLogger.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Info("Not downloading")
			continue
		}
		downloadID, err := show.AddDownload(con.db, res.Title, res.Magnetlink)
		if err != nil {
			resLogger.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Info("Adding download failed")
			continue
		}

		downloadPath := fmt.Sprintf("%s/%s", con.config.GetString("download_path"), downloadID)
		resLogger = resLogger.WithFields(logrus.Fields{
			"downloadid":   downloadID,
			"downloadpath": downloadPath,
		})

		resLogger.Info("Starting download")

		downloadResult := downloader.SyncDownload(res.Magnetlink, downloadPath)
		if !downloadResult.Completed {
			resLogger.WithFields(logrus.Fields{
				"error": downloadResult.Error.Error(),
			}).Warning("Download did not complete")

			show.SetDownloadFailed(con.db, res.Title)
			if con.config.GetBool("metrics_enabled") {
				failedDownloadsCounter.Inc()
			}
			continue
		}

		resLogger.Info("Download completed")
		if con.config.GetBool("metrics_enabled") {
			completedDownloadsCounter.Inc()
		}
		var largestFile *torrent.File
		var largest int64
		largest = 0
		for _, f := range downloadResult.Files {
			resLogger.WithFields(logrus.Fields{
				"filename": f.Path(),
			}).Debug("Found file")
			if f.Length() > largest {
				largest = f.Length()
				largestFile = f
			}
		}
		showPath := show.Path(con.config.GetString("shows_path"))
		targetDir := show.GetSeasonDir(res.Title, showPath)
		targetName := filepath.Join(targetDir, filepath.Base(largestFile.Path()))
		sourceName := filepath.Join(downloadPath, largestFile.Path())

		resLogger = resLogger.WithFields(logrus.Fields{
			"showpath":   showPath,
			"targetdir":  targetDir,
			"targetname": targetName,
			"sourcename": sourceName,
		})

		err = os.MkdirAll(filepath.Dir(targetName), 0777)
		if err != nil {
			resLogger.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Warning("Could not create directory")
			continue
		}

		err = Copy(sourceName, targetName)
		if err != nil {
			resLogger.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Warning("Could not copy file to target location")
			continue
		}
		os.RemoveAll(downloadPath)
		resLogger.Debug("Marking as downloaded")
		show.SetDownloaded(con.db, res.Title)
	}
}

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func validCodec(title string) bool {
	t := strings.ToLower(title)
	return strings.Contains(t, "264") ||
		strings.Contains(t, "265") ||
		strings.Contains(t, "hevc")
}

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gnur/go-piratebay"
	"github.com/gnur/golpje/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	if os.Getenv("GOLPJE_DAEMONIZE") != "" {
		// do this in a loop
		log.Info("doing nothing, this would do runonce in a loop")
	}

	runOnce()
}

func runOnce() {
	cfg, err := config.Load()
	if err != nil {
		log.WithField("error", err).Error("Loading configuration failed")
		return
	}
	for name, show := range cfg.Shows {
		if show.Active {
			showLog := log.WithField("show", name)
			show.Episodes = make(map[string]bool)
			showLog.Debug("starting work")
			files, err := cfg.Store.List(name, "")
			if err != nil {
				showLog.WithField("err", err).Warning("listing failed")
				continue
			}
			for _, file := range files {
				if looksLikeAVideo(file) {
					id, err := ExtractEpisodeID(file, show.Seasonal)
					if err == nil {
						show.Episodes[id] = true
					}
				}
			}
			pb := piratebay.New(cfg.Piratebayurl)
			torrents, err := pb.Search(name)
			if err != nil {
				showLog.WithField("err", err).Warning("Searching failed")
				continue
			}
			showLog.WithField("results", len(torrents)).Info("Found torrents")
			for _, result := range torrents {
				dl, id := ShouldDownload(show, result)
				showLog.WithFields(log.Fields{
					"episodeid":   id,
					"downloading": dl,
				}).Debug("Found result")
				if !dl {
					continue
				}
				dlDir := "/tmp/golpje-dl"
				os.RemoveAll(dlDir)
				os.MkdirAll(dlDir, 0777)
				dlResult := Download(result.MagnetLink, dlDir)
				if !dlResult.Completed {
					showLog.WithField("err", dlResult.Error.Error()).Warning("download did not complete")
					continue
				}

				for _, f := range dlResult.Files {
					if !looksLikeAVideo(f.Path()) {
						showLog.WithFields(log.Fields{
							"path": f.Path(),
						}).Debug("ignoring this file")
						continue
					}
					sourceName := filepath.Join(dlDir, f.Path())
					extension := filepath.Ext(sourceName)
					targetName := filepath.Join(FormatTargetDir(name, id, show.Seasonal), filepath.Base(f.Path()))
					if extension != ".mp4" && cfg.ConvertToMP4 {
						sourceConvert := strings.Replace(sourceName, extension, ".mp4", -1)
						showLog.WithFields(log.Fields{
							"src":    sourceName,
							"target": sourceConvert,
						}).Info("converting container")
						cmd := exec.Command("ffmpeg", "-i", sourceName, "-codec", "copy", sourceConvert)
						err := cmd.Run()
						if err != nil {
							showLog.WithFields(log.Fields{
								"src":    sourceName,
								"target": sourceConvert,
								"error":  err.Error(),
							}).Warning("Converting failed")
							continue
						}
						sourceName = sourceConvert
						targetName = strings.Replace(targetName, extension, ".mp4", -1)
					}
					showLog.WithFields(log.Fields{
						"src":    sourceName,
						"target": targetName,
					}).Info("starting transfer")
					in, err := os.Open(sourceName)
					if err != nil {
						showLog.WithField("err", err).Error("Could not open file for reading")
						continue
					}
					defer in.Close()

					err = cfg.Store.Write(targetName, in)
					if err != nil {
						showLog.WithField("err", err).Error("Could not transfer file to target dir")
						continue
					}

					show.Episodes[id] = true
				}
			}
		}
	}
}

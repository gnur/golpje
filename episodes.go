package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gnur/go-piratebay"
	"github.com/gnur/golpje/config"
	log "github.com/sirupsen/logrus"
)

// Episode is the struct that holds an episode
type Episode struct {
	ID          string
	Title       string
	Showid      string
	Added       int64
	Magnetlink  string
	Episodeid   string
	Downloaded  bool
	Downloading bool
}

var (
	episodeRegexp      = regexp.MustCompile(`.*?[sS]?([0-9]{1,2})[xeXE]([0-9]{1,2}).*`)
	episodeShortRegexp = regexp.MustCompile(`.*[^2-9]([0-9]{1})([0-3][0-9])[^0-9p].*`)
	dateRegexp         = regexp.MustCompile(`.*([0-9]{4})\.?\s?(\d\d)\.?\s?(\d\d).*`)
)

// ShouldDownload takes a show and torrent and returns whether this should be downloaded
func ShouldDownload(show *config.Show, res piratebay.Torrent) (bool, string) {

	l := log.WithFields(log.Fields{
		"show":  show.Regexp,
		"title": res.Title,
	})

	re := regexp.MustCompile("(?i)" + show.Regexp)
	if !re.MatchString(res.Title) {
		l.Info("result does not match show")
		return false, ""
	}

	id, err := ExtractEpisodeID(res.Title, show.Seasonal)
	if err != nil {
		l.WithField("err", err).Info("Could not extract episodeid")
		return false, ""
	}

	if res.Seeders < 10 {
		l.Debug("not downloading, not enough seeders")
		return false, id
	}
	if !validCodec(res.Title) {
		l.Debug("not downloading, no valid codec in title")
		return false, id
	}

	if _, ok := show.Episodes[id]; ok {
		l.Debug("Episode already in map")
		return false, id
	}

	return NewEnough(id, show.Minimal, show.Maxage, show.Seasonal), id
}

// NewEnough checks a episodeid for newenoughiness
func NewEnough(id string, minimal, maxage int64, seasonal bool) bool {
	if !seasonal {
		date, err := time.Parse("2006-01-02", id)
		if err != nil {
			log.Debug("could not parse date")
			return false
		}
		return time.Since(date) < 7*24*time.Hour
	}
	season, err := strconv.ParseInt(id[1:3], 10, 64)
	if err != nil {
		return false
	}
	return season >= minimal
}

// ExtractEpisodeID generates the episode id from a show title
func ExtractEpisodeID(title string, seasonal bool) (string, error) {
	var matches []string

	if !seasonal {
		matches = dateRegexp.FindStringSubmatch(title)
		if matches != nil && len(matches) == 4 {
			return strings.Join(matches[1:], "-"), nil
		}
	} else {
		matches = episodeRegexp.FindStringSubmatch(title)
		if matches != nil && len(matches) == 3 {
			season, err := strconv.ParseInt(matches[1], 10, 64)
			if err != nil {
				return "", errors.New("could not parse season")
			}
			episode, err := strconv.ParseInt(matches[2], 10, 64)
			if err != nil {
				return "", errors.New("could not parse episode")
			}
			return fmt.Sprintf("s%02de%02d", season, episode), nil
		}
		matches = episodeShortRegexp.FindStringSubmatch(title)
		if matches != nil && len(matches) == 3 {
			season, err := strconv.ParseInt(matches[1], 10, 64)
			if err != nil {
				return "", errors.New("could not parse season")
			}
			episode, err := strconv.ParseInt(matches[2], 10, 64)
			if err != nil {
				return "", errors.New("could not parse episode")
			}
			return fmt.Sprintf("s%02de%02d", season, episode), nil
		}
	}
	return "", errors.New("Could not extract episode id")
}

func validCodec(title string) bool {
	t := strings.ToLower(title)
	return strings.Contains(t, "264")
}

// FormatTargetDir returns the seasonal directory where episodes with given title are stored
func FormatTargetDir(showname, episodeid string, seasonal bool) string {

	var seasonDir string
	if !seasonal {
		seasonDir = episodeid[:7]
	} else {
		seasonDir = fmt.Sprintf("season %v", episodeid[1:3])
	}

	return filepath.Join(showname, seasonDir)
}

func looksLikeAVideo(s string) bool {
	parts := strings.Split(s, ".")
	if len(parts) < 2 {
		return false
	}
	ext := parts[len(parts)-1]
	return ext == "mp4" || ext == "mkv" || ext == "avi" || ext == "wmv"
}

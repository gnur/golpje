package episodes

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/asdine/storm"
	"github.com/google/uuid"
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

// All returns all events
func All(db *storm.DB) ([]Episode, error) {
	var episodes []Episode

	err := db.AllByIndex("Timestamp", &episodes, storm.Reverse())
	if err != nil {
		fmt.Println(err.Error())
	}
	return episodes, err
}

// New adds an episode and returns if it was successful
func New(db *storm.DB, title, showID, episodeID, magnetlink string, downloaded, downloading bool) (string, error) {

	u1 := uuid.New()

	e := Episode{
		ID:          u1.String(),
		Title:       title,
		Showid:      showID,
		Episodeid:   episodeID,
		Magnetlink:  magnetlink,
		Downloaded:  downloaded,
		Downloading: downloading,
	}
	err := db.Save(&e)
	return e.ID, err
}

// NewEnough checks a episodeid for newenoughiness
func NewEnough(db *storm.DB, id string, minimal int64, seasonal bool) bool {
	if !seasonal {
		year, err := strconv.ParseInt(id[0:4], 10, 64)
		if err != nil {
			return false
		}
		return year >= minimal
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

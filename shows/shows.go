package shows

import (
	"errors"
	"fmt"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/gnur/golpje/database"
	"github.com/gnur/golpje/episodes"
	"github.com/gnur/golpje/golpje"
	"github.com/google/uuid"
)

// Show a local alias of the protobuf Show
type Show struct {
	ID       string
	Name     string
	Regexp   string
	Active   bool
	Seasonal bool
	Minimal  int64
}

// New creates a new show
func New(name, regexp string, seasonal, active bool, minimal int64) (string, error) {
	var match Show

	err := database.Conn.One("Name", name, &match)
	if err != storm.ErrNotFound {
		return match.ID, errors.New("Show with this name already exists")
	}

	u1 := uuid.New()

	s := Show{
		ID:       u1.String(),
		Name:     name,
		Regexp:   regexp,
		Seasonal: seasonal,
		Active:   active,
		Minimal:  minimal,
	}
	err = database.Conn.Save(&s)
	return s.ID, nil
}

// All returns all shows
func All() ([]Show, error) {
	var matches []Show
	err := database.Conn.All(&matches)
	return matches, err
}

// GetFromID retrieve a show from an UUID
func GetFromID(uuid string) (Show, error) {

	var match Show

	err := database.Conn.One("ID", uuid, &match)
	if err != nil {
		return match, err
	}
	return match, nil
}

// GetFromName retrieve a show from an UUID
func GetFromName(name string) (Show, error) {
	var match Show

	err := database.Conn.One("name", name, &match)
	if err != nil {
		return match, err
	}
	return match, nil
}

// Delete removes a show from the database
func (s Show) Delete() error {
	return database.Conn.DeleteStruct(&s)
}

// Print provides a convenient way of pretty printing a show
func (s Show) Print() {
	fmt.Println("--------------")
	fmt.Println(s.ID, s.Name, s.Active)
}

// ToProtoShows returns a proto shows from a single show
func ToProtoShows(shows []Show) *golpje.ProtoShows {
	var resp golpje.ProtoShows
	for _, s := range shows {
		resp.Shows = append(resp.Shows, s.ToProto())
	}
	return &resp
}

// FromProto converts a proto message to a Show
func FromProto(in *golpje.ProtoShow) Show {
	return Show{
		ID:       in.ID,
		Name:     in.Name,
		Regexp:   in.Regexp,
		Active:   in.Active,
		Seasonal: in.Seasonal,
		Minimal:  in.Minimal,
	}
}

// ToProto converts a Show to a proto message
func (s Show) ToProto() *golpje.ProtoShow {
	sProto := golpje.ProtoShow{
		ID:       s.ID,
		Name:     s.Name,
		Regexp:   s.Regexp,
		Active:   s.Active,
		Seasonal: s.Seasonal,
		Minimal:  s.Minimal,
	}
	return &sProto
}

// ShouldDownload returns if the episode has not been downloaded yet and still should
func (s Show) ShouldDownload(title string) bool {

	episodeid, err := episodes.ExtractEpisodeID(title, s.Seasonal)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if !episodes.NewEnough(episodeid, s.Minimal, s.Seasonal) {
		return false
	}

	query := database.Conn.Select(q.Eq("Showid", s.ID), q.Eq("Episodeid", episodeid))

	var episodes []episodes.Episode
	err = query.Find(&episodes)

	if err != nil && err.Error() != "not found" {
		fmt.Println(err.Error())
		return false
	}

	for _, episode := range episodes {
		if episode.Downloading || episode.Downloaded {
			return false
		}
	}

	return true
}

// AddDownload adds an episode to a show
func (s Show) AddDownload(title, magnetlink string) (string, error) {
	fmt.Println("adding download")

	episodeID, err := episodes.ExtractEpisodeID(title, s.Seasonal)
	if err != nil {
		return "", err
	}

	return episodes.New(title, s.ID, episodeID, magnetlink, false, true)
}

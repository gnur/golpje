package config

import (
	"errors"
	"os"
	"time"

	"github.com/burntsushi/toml"
	"github.com/gnur/s3local"

	log "github.com/sirupsen/logrus"
)

// Cfg holds the config
type Cfg struct {
	Store          s3local.Store
	Searchinterval duration
	LogLevel       string
	Shows          map[string]*Show
	ConvertToMP4   bool
}

// Show holds all the information about a TV show
type Show struct {
	Seasonal bool
	Active   bool
	Maxage   int64
	Minimal  int64
	Regexp   string
	Episodes map[string]bool
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

// Load uses env GOLPJE_STORAGE to determine storage location
func Load() (*Cfg, error) {
	log.Debug("loading")
	var v Cfg
	var err error
	var store s3local.Store
	location := os.Getenv("GOLPJE_STORAGE")
	if location == "" {
		log.Info("using current directory")
		return nil, errors.New("not implemented yet")

	} else if location == "s3://" {
		log.Info("using s3")
		store, err = s3local.New(s3local.Config{
			Type: "s3",
			Settings: map[string]string{
				"host":            os.Getenv("S3_HOST"),
				"secretaccesskey": os.Getenv("S3_SECRET_ACCESS_KEY"),
				"accesskeyid":     os.Getenv("S3_ACCESS_KEY_ID"),
				"bucket":          os.Getenv("S3_BUCKET"),
			},
		})
		if err != nil {
			log.Error("Creating s3 store failed")
			return nil, err
		}
	} else {
		log.Info("using filesystem")
		store, _ = s3local.New(s3local.Config{
			Type: "local",
			Settings: map[string]string{
				"path": location,
			},
		})
	}

	v.Store = store
	log.Debug("loading golpje.toml")
	f, err := v.Store.Read("golpje.toml")
	if err != nil {
		log.WithField("err", err).Warning("could not read golpje.toml")
		return nil, err
	}
	log.WithField("config", string(f)).Debug("decoding golpje.toml")
	_, err = toml.Decode(string(f), &v)
	if err != nil {
		log.WithField("err", err).Warning("could not decode golpje.toml")
		return nil, err
	}
	log.Debug("loaded config")

	return &v, nil
}

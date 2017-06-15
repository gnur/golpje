package config

import "github.com/spf13/viper"

// Load parses env, files (and will parse flags) into a configuration
func Load() (*viper.Viper, error) {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvPrefix("golpje")

	v.SetConfigName("config")
	v.AddConfigPath("/etc/golpje/")
	v.AddConfigPath("$HOME/.golpje")
	v.AddConfigPath(".")
	v.ReadInConfig()
	loadDefaults(v)

	return v, nil
}

func loadDefaults(v *viper.Viper) {

	v.SetDefault("shows_path", "./shows/")
	v.SetDefault("database_file", "golpje.db")
	v.SetDefault("search_enabled", false)
}
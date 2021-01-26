package config

import (
	"strings"

	"github.com/spf13/viper"
)

// Database Config
type Database struct {
	URL string
}

// Configuration App vars
type Configuration struct {
	Port     string
	LogLevel string
	Database Database
}

func bindKeys() {
	viper.BindEnv("LogLevel", "LOG_LEVEL")
	viper.BindEnv("Port", "PORT")
	viper.BindEnv("Database.URL", "DATABASE_URL")
}

func setDefaults() {
	viper.SetDefault("LogLevel", "debug")
	viper.SetDefault("Port", "5001")
	viper.SetDefault("Database.URL", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
}

// New configuration
func New() (*Configuration, error) {
	conf := &Configuration{}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	bindKeys()

	setDefaults()

	viper.AutomaticEnv()

	err := viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

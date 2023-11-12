package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configuration Data for config object.
type CFGData struct {
	Port string
	Dsn  string
}

func NewCFGData() CFGData {

	viper.SetConfigName("config")

	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")

	viper.SetDefault("environments.dev.port", ":3000")
	viper.SetDefault("environments.dev.dsn", "mongodb://localhost:27017")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("err: %w", err))
	}

	return CFGData{
		Port: viper.GetString("environments.dev.port"),
		Dsn:  viper.GetString("environments.dev.dsn"),
	}
}

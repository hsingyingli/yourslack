package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB_USERNAME string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_TABLE    string `mapstructure:"DB_TABLE"`
	DB_URL      string `mapstructure:"DB_URL"`
	PORT        string `mapstructure:"PORT"`
	GIN_MODE    string `mapstructure:"GIN_MODE"`
}

func LoadConfig() (Config, error) {
	var config Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}

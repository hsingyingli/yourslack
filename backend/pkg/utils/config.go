package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_USERNAME            string        `mapstructure:"DB_USERNAME"`
	DB_PASSWORD            string        `mapstructure:"DB_PASSWORD"`
	DB_TABLE               string        `mapstructure:"DB_TABLE"`
	DB_URL                 string        `mapstructure:"DB_URL"`
	SYMMETRICKEY           string        `mapstructure:"SYMMETRICKEY"`
	ACCESS_TOKEN_DURATION  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	REFRESH_TOKEN_DURATION time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	COOKIE_MAXAGE          int           `mapstructure:"COOKIE_MAXAGE"`
	PORT                   string        `mapstructure:"PORT"`
	GIN_MODE               string        `mapstructure:"GIN_MODE"`
	ALLOW_ORIGIN           string        `mapstructure:"ALLOW_ORIGIN"`
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

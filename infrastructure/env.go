package infrastructure

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	DBName                 string `mapstructure:"DB_NAME"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerPort             string `mapstructure:"SERVER_PORT"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() Env {
	env := Env{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("App configuration not found")
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Can't load configuration file")
	}
	return env
}

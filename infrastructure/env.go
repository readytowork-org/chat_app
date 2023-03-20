package infrastructure

import (
	"log"

	"github.com/spf13/viper"
)

// Env has environment stored
type Env struct {
	ServerPort                   string `mapstructure:"ServerPort"`
	Environment                  string `mapstructure:"Environment"`
	JWT_ACCESS_SECRET            string `mapstructure:"JWT_ACCESS_SECRET"`
	JWT_REFRESH_SECRET           string `mapstructure:"JWT_REFRESH_SECRET"`
	JWT_ACCESS_TOKEN_EXPIRES_AT  int    `mapstructure:"JWT_ACCESS_TOKEN_EXPIRES_AT"`
	JWT_REFRESH_TOKEN_EXPIRES_AT int    `mapstructure:"JWT_REFRESH_TOKEN_EXPIRES_AT"`
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("☠️ Env config file not found: ", err)
		} else {
			log.Fatal("☠️ Env config file error: ", err)
		}
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	log.Printf("%+v \n", env)
	return env
}

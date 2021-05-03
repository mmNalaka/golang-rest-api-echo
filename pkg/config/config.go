package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DbHost       string `mapstructure:"DB_HOST"`
	DbPort       string `mapstructure:"DB_PORT"`
	DbName       string `mapstructure:"DB_NAME"`
	DbUser       string `mapstructure:"DB_USER"`
	DbPassword   string `mapstructure:"DB_PASSWORD"`
	Env          string `mapstructure:"ENV"`
	JWtExpiresIn string `mapstructure:"JWT_EXPIRES_IN"`
	JWtSecret    string `mapstructure:"JWT_SECRET"`
}

func New() *AppConfig {
	var cfg AppConfig

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No .env file, using environment variables", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Unable to decode config variables", err)
	}

	return &cfg
}

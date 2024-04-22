package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func init() {
	log.Println("Initializing configuration setup")
	env := os.Getenv("ENVIRONMENT")
	if env == "" || env == "DEVELOPMENT" {
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err != nil {
			log.Panicf("Error reading config file, %s", err)
		}
		viper.SetDefault("ENVIRONMENT", "DEVELOPMENT")
	}

	if env == "PRODUCTION" {
		viper.AutomaticEnv()
	}

}

type Config struct {
	SuperUserName   string
	SuperUserPass   string
	Environment     string
	Port            string
	Prefork         bool
	DBURL           string
	DBAPPURL        string
	ApiHost         string
	IsProduction    bool
	IsDevelopment   bool
	ApplicationName string
	AppDefaultURL   string
}

var config *Config

func LoadConfig() {
	config = &Config{
		SuperUserName:   viper.GetString("SUPERUSERNAME"),
		SuperUserPass:   viper.GetString("SUPERUSERPASS"),
		Environment:     viper.GetString("ENVIRONMENT"),
		Port:            viper.GetString("PORT"),
		Prefork:         viper.GetBool("PREFORK"),
		DBURL:           viper.GetString("DATABASE_URL"),
		DBAPPURL:        viper.GetString("DATABASE_APP_URL"),
		ApiHost:         viper.GetString("API_ENDPOINT"),
		IsProduction:    viper.GetString("ENVIRONMENT") == "PRODUCTION",
		IsDevelopment:   viper.GetString("ENVIRONMENT") != "PRODUCTION",
		ApplicationName: viper.GetString("APPLICATION_NAME"),
		AppDefaultURL:   viper.GetString("APP_DEFAULT_URL"),
	}
}

func GetConfig() *Config {
	return config
}

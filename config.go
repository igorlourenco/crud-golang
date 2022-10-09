package main

import (
	"log"

	"github.com/spf13/viper"
)

/*
 * this is the struct that will contain the allowed configurations
 * port number and postgresql connection string
 */
type Config struct {
	Port            string `mapstructure:"port"`
	ConnectionSting string `mapstructure:"connection_string"`
}

// this is the global variable that will be used to access the configurations
var AppConfig *Config

func LoadAppConfig() {

	// using viper to load the configurations from the config.json file
	log.Println("Loading server configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&AppConfig)

	if err != nil {
		log.Fatal(err)
	}
}

package config

import (
	"github.com/spf13/viper"
	model "github.com/MufidJamaluddin/simplewebgo/model"
)

var appConfig model.ConfigurationModel

func init() {
	readConfig()
}

// GetConfig : Mendapatkan Konfigurasi
func GetConfig() model.ConfigurationModel {
	return appConfig
}

func readConfig() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	//viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&appConfig)
	if err != nil {
		panic(err)
	}
}
package config

import (
	"log"
	
	

	"github.com/spf13/viper"	
)

var (

	JWTKey  = ""
)

type Configuration struct {
	Port    string                    `mapstructure:"Port"`
	Database struct {  
			Driver             string `mapstructure:"Driver"`
			Host               string `mapstructure:"Host"`
			Name               string `mapstructure:"Name"`
			Port               string `mapstructure:"Port"`
			Username           string `mapstructure:"Username"`
			Password           string `mapstructure:"Password"`
			JwtKey             string `mapstructure:"JwtKey"`
	}       `mapstructure:"database"`
} 


var appConfig *Configuration 

func GetConfiguration() *Configuration {
	return InitConfiguration()
} 

func InitConfiguration() *Configuration {
	app := Configuration{}
	 
		viper.AddConfigPath(".")
		viper.SetConfigName("app")
		viper.SetConfigType("env") 
		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil 
		} 
		err = viper.Unmarshal(&app)
		if err != nil {
			log.Println("error parse config : ", err.Error())
		}

		JWTKey    = app.Database.JwtKey

		return &app
}
package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

type GeneralConfig struct {
	host  string `env:"MACSERVER_HOST"`
	port  int `env:"MACSERVER_PORT"`
	debug bool `env:"MACSERVER_DEBUG"`

	dbUsername string `env:"MACSERVER_DB_USERNAME"`
	dbPassword string `env:"MACSERVER_DB_PASSWORD"`
	dbHost string `env:"MACSERVER_DB_HOST"`
	dbPort int `env:"MACSERVER_DB_PORT"`
	dbName string `env:"MACSERVER_DB_NAME"`
}

var config GeneralConfig

func init() {
	err := godotenv.Load("../Config.env")

	if err != nil {

		log.Println("Error loading .env file")
		config.LoadDefaultConfig()

	} else {
		log.Println("File loaded!")

		if config.LoadEnvConfig() != nil{

			log.Println("An error has occurred while I was reading the environment variables. Loading default configuration...")
			config.LoadDefaultConfig()

		} else {
			log.Println("Environment Configuration gotten successfully.")
			log.Println("Your configuration is: ",config)
		}
	}
}

func Debug() bool {
	return config.debug
}

func GetUrlDatabase() string {
	return config.GetDatabaseUrl()
}

func GetUrlHttp() string {
	return config.GetHttpUrl()
}

func (this *GeneralConfig) GetDatabaseUrl () string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", this.dbUsername, this.dbPassword, this.dbHost, this.dbPort, this.dbName)
}

func (this *GeneralConfig) GetHttpUrl () string {
	return fmt.Sprintf("%s:%d", this.host, this.port)
}

func Application() string {
	return "./www/application/*.html"
}

func ApplicationError() string {
	return "./www/application/error/404.html"
}

func Assets() string {
	return "./www/application/assets"
}

func (this *GeneralConfig) LoadDefaultConfig(){
	config.debug = true
	config.host = "localhost"
	config.port = 8000
	config.dbUsername = "root"
	config.dbPassword = ""
	config.dbHost = "localhost"
	config.dbPort = 3306
	config.dbName = "test"
}

func (this *GeneralConfig) LoadEnvConfig() error{

	SetPrefix("MACSERVER")

	config.debug = GetBoolFromEnv("DEBUG")
	config.host = GetStringFromEnv("HOST")
	config.port, _ = GetIntFromEnv("PORT")
	config.dbUsername = GetStringFromEnv("DB_USERNAME")
	config.dbPassword = GetStringFromEnv("DB_PASSWORD")
	config.dbHost = GetStringFromEnv("DB_HOST")
	config.dbPort, _ = GetIntFromEnv("DB_PORT")
	config.dbName = GetStringFromEnv("DB_NAME")

	if config.dbPort == 0{
		return errors.New("no correct values")
	}

	return nil
}
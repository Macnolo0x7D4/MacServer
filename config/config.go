package config

import (
	"fmt"
)

type DatabaseConfig struct {
	username string
	password string
	host     string
	port     int
	database string
}

type ServerConfig struct {
	host  string
	port  int
	debug bool
}

type Config interface {
	url() string
}

var database *DatabaseConfig
var server *ServerConfig

func init() {
	database = &DatabaseConfig{}
	/*
		database.username = gonv.GetStringEnv("MACSERVER_DATABASE_USERNAME", "root")
		database.password = gonv.GetStringEnv("MACSERVER_DATABASE_PASSWORD", " ")
		database.host = gonv.GetStringEnv("MACSERVER_DATABASE_HOST", "localhost")
		database.port = gonv.GetIntEnv("MACSERVER_DATABASE_PORT", 3306)
		database.database = gonv.GetStringEnv("MACSERVER_DATABASE_NAME", "test")*/

	database.username = "root"
	database.password = " "
	database.host = "localhost"
	database.port = 3306
	database.database = "test"

	server = &ServerConfig{}
	/*server.host = gonv.GetStringEnv("MACSERVER_WEB_HOST", "localhost")
	server.port = gonv.GetIntEnv("MACSERVER_WEB_PORT", 8000)
	server.debug = gonv.GetBoolEnv("MACSERVER_MODE", true)*/
	server.host = "localhost"
	server.port = 8000
	server.debug = true
}

func Debug() bool {
	return server.debug
}
func GetUrlDatabase() string {
	return database.url()
}

func GetUrlWebserver() string {
	return server.url()
}

func (this *DatabaseConfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", this.username, this.password, this.host, this.port, this.database)
}

func (this *ServerConfig) url() string {
	return fmt.Sprintf("%s:%d", this.host, this.port)
}

func ServerPort() int {
	return server.port
}

func Application() string {
	return "../www/**/*.html"
}

func ApplicationError() string {
	return "../www/application/error/404.html"
}

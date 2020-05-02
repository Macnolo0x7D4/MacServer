package database

import (
	"../config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var Database *gorm.DB

func CreateConnection() {
	if connection, err := gorm.Open("mysql", config.GetUrlDatabase()); err != nil {
		panic(err)
	} else {
		log.Println("MySQL Connection successfully")
		Database = connection
	}
}

func CloseConnection() {
	err := Database.Close()

	if err != nil{
		panic(err)
	}

}

func CreateTable() {
	Database.DropTableIfExists(&User{})
	Database.CreateTable(&User{})
}

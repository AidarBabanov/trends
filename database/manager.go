package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	host := os.Getenv("db_host")
	port := os.Getenv("db_port")
	database := os.Getenv("db_name")
	dialect := os.Getenv("db_type")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC", username, password, host, port, database)

	log.Println("Connecting to database.")
	log.Println(fmt.Sprintf("Database credentials: \n"+
		"	Host: %s\n"+
		"	Port: %s\n"+
		"	Database name: %s\n"+
		"	Username: %s\n"+
		"	Password: %s", host, port, database, username, password))
	var err error
	DB, err = gorm.Open(dialect, dataSource)
	if err == nil {
		log.Println("Connected successfully.")
	} else {
		log.Println("Connection failed.")
	}
	return DB, err
}

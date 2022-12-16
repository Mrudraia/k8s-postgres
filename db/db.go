package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func DbURL() string {

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", "spuser", "SPuser96", "project", "postgresdb", "5432")
}

func Open() {
	DB, err = gorm.Open(postgres.Open(DbURL()), &gorm.Config{})

	if err != nil {
		log.Fatal("Problem opening database: " + err.Error())
	}
}

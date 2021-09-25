package models

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_CONNECTION")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	setupModels()
}

func setupModels() {
	DB.AutoMigrate(Bulb{})
	DB.AutoMigrate(Link{})
}

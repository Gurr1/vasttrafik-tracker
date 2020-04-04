package service

import (
	"github.com/jinzhu/gorm"
	"engsmyre.xyz/vasttrafik-tracker/database"
)

var Database gorm.DB

func Connect() {
	database.Connect()
}



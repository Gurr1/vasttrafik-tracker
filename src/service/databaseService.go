package service

import (
	"github.com/gurr1/vasttrafik-tracker/database"
	"log"
)


func Connect() {
	err := database.Connect()
	if err != nil {
		log.Fatal("Could not connect to database")
	}
	database.CreateAuthentication("token")
	database.GetActiveAuthentication()
}



package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

func Connect() (*gorm.DB, error) {
	host := "localhost"
	password := "password"
	user := "user"
	port := 5432
	db, err := gorm.Open("postres",
		fmt.Sprintf("host=%s, user=%s, port=%d, dname=database.db, password=%s", host, user, port, password))
	if err != nil {
		log.Fatal("Could not connect to database, aborting launch")
	}
	defer db.Close()		// This should not be done as such, should perhaps be rewritten
	return db, nil
}

func Migrate() {
	//TODO
}
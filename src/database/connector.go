package database

import (
	"fmt"
	"github.com/gurr1/vasttrafik-tracker/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var Database *gorm.DB

func Connect() error {
	host := "localhost"
	password := "password"
	user := "user"
	port := 5432
	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s user=%s port=%d dbname=database password=%s sslmode=disable", host, user, port, password))
	if err != nil {
		fmt.Println(err)
		log.Fatal("Could not connect to database, aborting launch")
	}
	//defer db.Close()		// This should not be done as such, should perhaps be rewritten
	db.AutoMigrate(&entity.AuthenticationDetails{})
	//db.Create(&entity.AuthenticationDetails{ID:uuid.New().String(), Token:"hej"})
	Database = db
	return nil
}

func Migrate() {
	//TODO
}
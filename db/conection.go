package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=pane password=1234 dbname=panedb port=8000" 

var DB *gorm.DB

func ConectionDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN),&gorm.Config{})
	if err != nil  {
		log.Fatal(err)
	}else{
		log.Println("DB connected 👍👍👍")
	}

}
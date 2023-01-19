package connection

import (
	"github.com/rakesh1602/GoLang/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connection() *gorm.DB {
	url := os.Getenv("url")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Println("Connection to database is failed !!")
		log.Fatal(err)
	}

	if db != nil {
		log.Println("Connected to database successfully.")
	}

	db.AutoMigrate(model.Employee{}, model.Department{})

	return db
}

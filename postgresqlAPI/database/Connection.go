package database

import (
	"github.com/rakesh1602/golang-postgres-book/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

//Connection using go-pg/pg
/*func Connect() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("userName"),
		Password: os.Getenv("password"),
		Addr:     os.Getenv("address"),
		Database: os.Getenv("database"),
	}

	//Most important thing
	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Fatal("Failed to connect to database")
		os.Exit(100)
	}
	log.Println("Connected successfully")

	err := createSchema(db)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func createSchema(db *pg.DB) interface{} {
	models := []interface{}{
		(*model.Book)(nil),
	}

	for _, mod := range models {
		err := db.Model(mod).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			log.Fatal(err)
		}
		log.Print("Created table in the database")

	}
	return nil

}*/

//var dsn = "host=localhost user=postgres password=1234 dbname=golang port=5432 sslmode=disable"

func Connection() *gorm.DB {
	url := os.Getenv("url")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if db != nil {
		log.Println("Connected successfully")
	}

	db.AutoMigrate(model.Book{})
	return db
}

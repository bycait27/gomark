package database

import (
	"gomark/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// load env variables
	dsn := "host=localhost user=caitlinash password=password dbname=gomark port=5432 sslmode=disable"

	// connect to postgresdatabase
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// aut-migrate the schema
	db.AutoMigrate(&models.Bookmark{})

	DB = db
	log.Println("Database connected successfully")
}

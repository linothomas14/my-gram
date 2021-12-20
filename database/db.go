package database

import (
	"fmt"
	"log"
	"my-gram/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func StartDB() {
	if os.Getenv("APP_ENV") != "production" {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	var (
		host     = os.Getenv("DB_HOST")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASS")
		dbPort   = os.Getenv("DB_PORT")
		dbName   = os.Getenv("DB_NAME")
	)

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbName, dbPort)
	if os.Getenv("APP_ENV") == "production" {
		config = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require ", host, user, password, dbName, dbPort)
	}
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to database", err)
	}

	fmt.Println("success connect to database")
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}

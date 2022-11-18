package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/shivammahatha/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {

var err  error

	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DbUser, DbPassword, DbName)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ")
	}

	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.Apipost{})
}
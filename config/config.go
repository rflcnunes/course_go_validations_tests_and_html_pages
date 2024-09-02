package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rflcnunes/course_validations_tests_and_html_pages/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func setupEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}
}

func setupDB() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dbURL))
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	DB.AutoMigrate(&models.Student{})

	fmt.Println("Database connection established successfully")
}

func Setup() {
	setupEnv()

	setupDB()
}

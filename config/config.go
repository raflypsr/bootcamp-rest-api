package config

import (
	"fp-super-bootcamp-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectDataBase() *gorm.DB {
	var db *gorm.DB

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	// production
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable"
	dbGorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db = dbGorm

	db.AutoMigrate(&models.Users{}, &models.Books{}, &models.Comments{}, &models.Likes{}, &models.Profile{}, &models.Reviews{})

	return db

}

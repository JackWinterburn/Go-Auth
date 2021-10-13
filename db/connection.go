package db

import (
	"os"

	"github.com/JackWinterburn/go-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(postgres.Open(os.Getenv("dsn")), &gorm.Config{})
	if err != nil {
		panic("Database connection failed!")
	}

	DB = conn

	if err := conn.AutoMigrate(&models.User{}); err != nil {
		panic("Could not migrate database models")
	}
}

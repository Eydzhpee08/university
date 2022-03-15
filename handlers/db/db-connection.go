package db

import (
	"fmt"
	"github.com/Eydzhpee08/university/handlers/models"
	"gorm.io/driver/postgres"
	"log"

	"gorm.io/gorm"
)

var Database *gorm.DB

var err error

//database migrate.
func DataMigration() {

	const (
		DB_USERNAME = "postgres"
		DB_HOST     = "localhost"
		DB_PORT     = "5435"
		DB_PASSWORD = "postgres"
		DB_NAME     = "univer"
		DB_SSLMODE  = "disable"
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		DB_HOST, DB_PORT, DB_USERNAME, DB_NAME, DB_PASSWORD, DB_SSLMODE)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("ERROR failed to connect database, err:", err)
		return
	}

	Database.AutoMigrate(
		&models.Employee{},
		&models.Sign{},
		&models.Docx{})
}

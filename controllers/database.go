// controllers/database.go

package controllers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() *gorm.DB {
	godotenv.Load()
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)
	database, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Error al conectar con la base de datos MySQL")
	}
	db = database
	return db
}

func GetDB() *gorm.DB {
	return db
}

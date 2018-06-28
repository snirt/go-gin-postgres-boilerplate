package conn

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// CreateDatabaseConnection - creating a DB connection
func CreateDatabaseConnection() (*gorm.DB, error) {
	HOST := getEnv("DB_HOST", "localhost")
	PORT := getEnv("DB_PORT", "5432")
	USER := getEnv("DB_USER", "postgres")
	DATABASE := getEnv("DB_DATABASE", "postgres")
	PASSWORD := getEnv("DB_PASSWORD", "agent007")

	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			HOST, PORT, USER, DATABASE, PASSWORD,
		),
	)
	if err != nil {
		panic(err.Error())
	}
	db.Exec("SET search_path TO golang")
	return db, err
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

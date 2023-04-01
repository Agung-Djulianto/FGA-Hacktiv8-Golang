package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ReadDB() *gorm.DB {
	const (
		host         = "localhost"
		user         = "postgres"
		password     = "admin"
		databasePort = "5432"
		databaseName = "belajar-gorm"
	)

	fix := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, databaseName, databasePort)

	db, err := gorm.Open(postgres.Open(fix), &gorm.Config{})

	if err != nil {
		log.Fatal("gagal tersambung kedatabase:", err)
	}

	return db
}

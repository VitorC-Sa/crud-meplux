package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() error {
	dsn := "host=localhost user=postgres password=123456 dbname=CRUD-MEPLUX port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return err
}

func GetConnection() *gorm.DB {
	return db
}

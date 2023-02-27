package utils

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=root dbname=demo port=5432 sslmode=disable"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		//panic("failed to connect database")
		return nil, err
	} else {
		fmt.Println("connected to database")
		return Db, nil
	}
}

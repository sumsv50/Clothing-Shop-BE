package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	dsn := "host=b4gdi69qs6nnyap9lgt2-postgresql.services.clever-cloud.com user=uz23oabarbdsfxmix2an password=PCdDdKnhrZH84uFHbm7pB8o3IK0COi dbname=b4gdi69qs6nnyap9lgt2 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect to DB successfully!")
	return db
}

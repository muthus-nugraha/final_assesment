package config

import (
	"final_assignment/app/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dbuser := "muthus"
	dbpass := "P@ssw0rd"
	dbhost := "localhost"
	dbname := "final_project"
	dbport := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbhost, dbuser, dbpass, dbname, dbport)
	db, errorDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Connecting DB Failed")
	}

	return db
}

func Disconnect(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Disconect DB Failed")
	}
	dbSQL.Close()
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.SocialMedia{}, &models.Comment{})
	fmt.Println("Migration DB Complated")
}

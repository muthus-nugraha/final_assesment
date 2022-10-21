package main

import (
	"final_assignment/app/routers"
	"final_assignment/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.Connect()
)

func main() {
	config.Migration(db)
	defer config.Disconnect(db)
	routers.InitRouter()
}

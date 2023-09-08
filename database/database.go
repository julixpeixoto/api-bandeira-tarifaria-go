package database

import (
	"api-bandeira-tarifaria-go/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=123456 dbname=flags port=15432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Migrate() {
	db := connect()
	db.AutoMigrate(&models.Flag{})
}

func Insert(flags []models.Flag) {
	db := connect()
	db.Create(&flags)
}

func GetAll() []models.Flag {
	db := connect()
	var flags []models.Flag
	result := db.Order("year desc, mount_num desc").Find(&flags)

	if result.Error != nil {
		return nil
	}

	return flags
}

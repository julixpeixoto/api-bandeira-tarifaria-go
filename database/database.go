package database

import (
	"api-bandeira-tarifaria-go/models"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connect() *gorm.DB {
	host := viper.GetString("DATABASE_HOST")
	user := viper.GetString("DATABASE_USER")
	password := viper.GetString("DATABASE_PASSWORD")
	name := viper.GetString("DATABASE_NAME")
	port := viper.GetString("DATABASE_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, name, port)
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

package config

import (
	"2hf/models"
	"2hf/utils"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	// "gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	environment := utils.Getenv("ENVIRONMENT", "development")
	if environment == "production" {
		username := os.Getenv("DATABASE_USERNAME")
		password := os.Getenv("DATABASE_PASSWORD")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")
		database := os.Getenv("DATABASE_NAME")
		// production
		dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=require"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Vocation{})

		// db.AutoMigrate(&models.Cart{}, &models.User{}, &models.User{})
		// db.AutoMigrate(&models.Cart{}, &models.StructShop{}, &models.User{})
		// db.AutoMigrate(&models.Transaction{}, &models.User{}, &models.User{})

		return db
	} else {

		username := "root"
		password := ""
		host := "tcp(192.168.217.88:3306)"
		database := "db_2hf"

		dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)
		// dsn := "root:root@tcp(192.168.217.88:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err.Error())
		}

		db.AutoMigrate(&models.User{}, &models.Vocation{}, &models.Payment{}, &models.Advertise{})

		return db
	}
}

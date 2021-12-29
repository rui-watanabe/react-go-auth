package database

import (
	"react-go-auth2/config"
	"react-go-auth2/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := config.Config.DBUser + ":" + config.Config.DBPass + "@tcp(" + config.Config.DBHost + ")/" + config.Config.DBName + "?charset=utf8&parseTime=True&loc=Local"
	// dsn = "root@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}

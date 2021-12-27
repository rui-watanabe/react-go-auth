package database

import (
	"react-go-auth2/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	dsn := config.Config.DBUser + ":" + config.Config.DBPass + "@tcp(" + config.Config.DBHost + ")/" + config.Config.DBName + "?charset=utf8&parseTime=True&loc=Local"
	// dsn = "root@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
}

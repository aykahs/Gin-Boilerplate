package mysql

import (
	"fmt"

	"github.com/aykahs/Gin-Boilerplate/configs"
	"github.com/aykahs/Gin-Boilerplate/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(config *configs.Mysql) *gorm.DB {

	// logger := logger.LogrusLogger

	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DataBase,
	)

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               address,
		DefaultStringSize: 256, // default size for string fields
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to MySQL")
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic("Failed to auto-migrate database schema")
	}

	DB = db

	return db

}
func GetDB() *gorm.DB {
	return DB
}

package utils

import (
	"github.com/testdev-learning/DataHub/backend/config"
	"github.com/testdev-learning/DataHub/backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	if config.DBType == "sqlite" {
		DB, err = gorm.Open(sqlite.Open("testdev.db"), &gorm.Config{})
	} else {
		dsn := config.DBUser + ":" + config.DBPassword + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	DB.AutoMigrate(
		&models.User{},
		&models.TestCase{},
		&models.TestReport{},
	)
}
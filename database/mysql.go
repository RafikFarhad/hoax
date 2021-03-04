package database

import (
	"fmt"
	"github.com/RafikFarhad/hoax/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMySqlDb(config *config.HoaxConfig) (*gorm.DB, error) {
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	var dbLogger logger.Interface
	if config.Logger != nil {
		dbLogger = *config.Logger
	}
	return gorm.Open(mysql.Open(connectionStr), &gorm.Config{
		Logger: dbLogger,
	})
}

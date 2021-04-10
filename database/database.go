package database

import (
	"errors"
	"fmt"
	"github.com/RafikFarhad/hoax/config"
	logger2 "github.com/RafikFarhad/hoax/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var AppDb *gorm.DB

func InitDatabase() error {
	appConfig := config.AppConfig
	var err error
	var connectionStr string
	switch appConfig.Db {
	case "mysql":
		connectionStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", appConfig.DbUser, appConfig.DbPassword, appConfig.DbHost, appConfig.DbPort, appConfig.DbName)
		break
	default:
		return errors.New("db of type " + appConfig.Db + " not yet supported")
	}

	var dbLogger logger.Interface
	if appConfig.DbLog {
		dbLogger = logger.New(
			logger2.GlobalLogger,
			logger.Config{
				LogLevel: logger.Info, // default db log info
				Colorful: true,
			})
	}
	AppDb, err = gorm.Open(mysql.Open(connectionStr), &gorm.Config{
		Logger: dbLogger,
	})
	return err
}

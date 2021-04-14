package database

import (
	"errors"
	"fmt"
	"github.com/RafikFarhad/hoax/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var AppDb *gorm.DB

func InitDatabase() error {
	appConfig := config.AppConfig
	var err error

	gormConfig := &gorm.Config{}
	if appConfig.DbConfig.Log {
		gormConfig.Logger = logger.New(
			&log.Logger,
			logger.Config{
				LogLevel: logger.Info,
				Colorful: false,
			})
	}
	switch appConfig.DbConfig.Agent {
	case "":
		return nil
	case "mysql":
		connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", appConfig.DbConfig.User, appConfig.DbConfig.Password, appConfig.DbConfig.Host, appConfig.DbConfig.Port, appConfig.DbConfig.Name)
		AppDb, err = gorm.Open(mysql.Open(connectionStr), gormConfig)
		break
	case "sqlite":
		connectionStr := fmt.Sprintf("%s", appConfig.DbConfig.Path)
		AppDb, err = gorm.Open(sqlite.Open(connectionStr), gormConfig)
		break
	default:
		return errors.New("db of type " + appConfig.DbConfig.Agent + " not yet supported")
	}
	return err
}

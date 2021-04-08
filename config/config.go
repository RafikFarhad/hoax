package config

import (
	"gopkg.in/ini.v1"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type HoaxConfig struct {
	Env     string `ini:"env"`
	Prefork bool   `ini:"prefork"`
	AppUrl  string `ini:"app_url"`
	Address string
	Debug   bool `ini:"debug"`

	Db         string `ini:"db"`
	DbHost     string `ini:"db_url"`
	DbPort     string `ini:"db_port"`
	DbName     string `ini:"db_name"`
	DbUser     string `ini:"db_user"`
	DbPassword string `ini:"db_password"`

	JwtSecret string `ini:"jwt_secret"`

	Logger *logger.Interface
}

var AppConfig *HoaxConfig

func ParseConfig(hostAddress string, configFile string) error {
	iniData, err := ini.Load(configFile)
	if err != nil {
		return err
	}

	AppConfig = &HoaxConfig{}

	_ = iniData.MapTo(AppConfig)
	_ = iniData.Section("deploy").MapTo(AppConfig)
	_ = iniData.Section("database").MapTo(AppConfig)
	_ = iniData.Section("auth").MapTo(AppConfig)

	AppConfig.Address = hostAddress

	// Prepare a app side logger
	if iniData.Section("log").Key("enabled").MustBool(false) {
		appLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.Lshortfile),
			logger.Config{
				LogLevel: logger.Info, // TODO:: <- need to handle config.ini value
				Colorful: true,
			})
		AppConfig.Logger = &appLogger
	}
	return nil
}

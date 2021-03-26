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

	DbHost     string `ini:"db_url"`
	DbPort     string `ini:"db_port"`
	DbName     string `ini:"db_name"`
	DbUser     string `ini:"db_user"`
	DbPassword string `ini:"db_password"`

	JwtSecret string `ini:"jwt_secret"`

	Logger *logger.Interface
}

func ParseConfig(hostAddress string, configFile string) (*HoaxConfig, error) {
	iniData, err := ini.Load(configFile)
	if err != nil {
		return nil, err
	}
	config := &HoaxConfig{}

	_ = iniData.MapTo(config)
	_ = iniData.Section("deploy").MapTo(config)
	_ = iniData.Section("database").MapTo(config)
	_ = iniData.Section("auth").MapTo(config)

	config.Address = hostAddress

	// Prepare a app side logger
	if iniData.Section("log").Key("enabled").MustBool(false) {
		appLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.Lshortfile),
			logger.Config{
				LogLevel: logger.Info, // TODO:: <- need to handle config.ini value
				Colorful: true,
			})
		config.Logger = &appLogger
	}

	//fmt.Printf("Config: %+v\n", config)
	return config, nil
}

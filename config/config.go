package config

import (
	"gopkg.in/ini.v1"
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
	DbLog      bool   `ini:"db_log"`

	JwtSecret string `ini:"jwt_secret"`

	SwaggerEnabled bool `ini:"enable_swagger"`
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
	_ = iniData.Section("utils").MapTo(AppConfig)

	AppConfig.Address = hostAddress

	return nil
}

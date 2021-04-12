package config

import (
	"gopkg.in/ini.v1"
)

type HoaxConfig struct {
	Env            string `ini:"env"`
	Prefork        bool   `ini:"prefork"`
	AppUrl         string `ini:"app_url"`
	Address        string
	Debug          bool   `ini:"debug"`
	JwtSecret      string `ini:"jwt_secret"`
	SwaggerEnabled bool   `ini:"enable_swagger"`

	DbConfig    *DbConfig
	CacheConfig *CacheConfig
}

var AppConfig *HoaxConfig

func ParseConfig(hostAddress string, configFile string) error {
	iniData, err := ini.Load(configFile)
	if err != nil {
		return err
	}

	AppConfig = &HoaxConfig{
		CacheConfig: &CacheConfig{},
		DbConfig:    &DbConfig{},
	}

	_ = iniData.MapTo(AppConfig)
	_ = iniData.Section("deploy").MapTo(AppConfig)
	_ = iniData.Section("auth").MapTo(AppConfig)
	_ = iniData.Section("utils").MapTo(AppConfig)
	_ = iniData.Section("database").MapTo(AppConfig.DbConfig)
	_ = iniData.Section("cache").MapTo(AppConfig.CacheConfig)

	AppConfig.Address = hostAddress
	return nil
}

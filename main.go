package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/RafikFarhad/hoax/cache"
	"github.com/RafikFarhad/hoax/config"
	"github.com/RafikFarhad/hoax/database"
	"github.com/RafikFarhad/hoax/database/model"
	_ "github.com/RafikFarhad/hoax/docs"
	"github.com/RafikFarhad/hoax/http"
	"github.com/RafikFarhad/hoax/logger"
	"github.com/RafikFarhad/hoax/routes"
)

var (
	configFile    string
	hostAddress   string
	autoMigration bool
)

// @title Hoax
// @version 1.0
// @description A ready eco-system to build web app faster on the go
// @contact.name Hoax Support
// @contact.email rafikfarhad@gmail.com
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @host localhost:3000
// @BasePath /
func main() {
	// Parse cmd args
	parseInputArgs()

	// Parse .env
	err := config.ParseConfig(hostAddress, configFile)
	//fmt.Printf("[DEBUG] %+v", config.AppConfig)

	appConfig := config.AppConfig
	if err != nil {
		panic(fmt.Sprintf("[INI] %+v", err))
	}

	// Create app
	if err := createApp(appConfig); err != nil {
		panic(err)
	}

	// Start Http Server
	err = http.AppHttp.Listen(appConfig.Address)
	if err != nil {
		panic(err)
	}
}

func parseInputArgs() {
	var host string
	var port int
	// Parse cmd args
	flag.StringVar(&host, "h", "127.0.0.1", "host")
	flag.IntVar(&port, "p", 3000, "port")
	flag.StringVar(&configFile, "c", "config.ini", "config file")
	flag.BoolVar(&autoMigration, "m", false, "run auto migration")
	flag.Parse()
	hostAddress = fmt.Sprintf("%s:%d", host, port)
}

func createApp(config *config.HoaxConfig) error {

	// global logger setup
	if err := logger.CreateLogger(); err != nil {
		fmt.Println("logger init error")
		return err
	}

	// http server setup
	if err := http.CreateHTTPServer(); err != nil {
		fmt.Println("HTTP init error")
		return err
	}
	routes.InitRoutes()

	// database setup
	if err := database.InitDatabase(); err != nil {
		fmt.Println("DB init error")
		return err
	}
	if autoMigration {
		if err := model.RunAutoMigrate(); err != nil {
			return errors.New("auto migration failed")
		}
	}

	// cache setup
	if err := cache.InitCache(); err != nil {
		fmt.Println("cache init error")
		return err
	}
	return nil
}

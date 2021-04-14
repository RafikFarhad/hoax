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
	// parse cmd args
	parseInputArgs()

	// parse .ini
	if err := config.ParseConfig(hostAddress, configFile); err != nil {
		panic(fmt.Sprintf("ini parsing failed :: %+v", err))
	}
	appConfig := config.AppConfig
	//fmt.Printf("[DEBUG] %+v", config.AppConfig)

	// create app
	if err := createApp(); err != nil {
		panic(err)
	}

	// start Http Server
	err := http.AppHttp.Listen(appConfig.Address)
	if err != nil {
		panic(err)
	}
}

func parseInputArgs() {
	var host string
	var port int
	// parse command line args
	flag.StringVar(&host, "h", "127.0.0.1", "host")
	flag.IntVar(&port, "p", 3000, "port")
	flag.StringVar(&configFile, "c", "config.ini", "config file")
	flag.BoolVar(&autoMigration, "m", false, "run auto migration")
	flag.Parse()
	hostAddress = fmt.Sprintf("%s:%d", host, port)
}

func createApp() error {
	// global logger setup
	if err := logger.InitLog(); err != nil {
		fmt.Println("logger init error")
		return err
	}
	// http server setup
	if err := http.CreateHTTPServer(); err != nil {
		return err
	}
	routes.InitRoutes()
	// database setup
	if err := database.InitDatabase(); err != nil {
		return err
	}
	if autoMigration {
		if err := model.RunAutoMigrate(); err != nil {
			return errors.New("auto migration failed")
		}
	}
	// cache setup
	if err := cache.InitCache(); err != nil {
		return err
	}
	return nil
}

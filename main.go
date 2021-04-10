package main

import (
	"flag"
	"fmt"
	"github.com/RafikFarhad/hoax/config"
	"github.com/RafikFarhad/hoax/database"
	_ "github.com/RafikFarhad/hoax/docs"
	"github.com/RafikFarhad/hoax/http"
	"github.com/RafikFarhad/hoax/logger"
	"github.com/RafikFarhad/hoax/routes"
)

var (
	configFile  string
	hostAddress string
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
	flag.StringVar(&host, "h", "127.0.0.1", "Host")
	flag.IntVar(&port, "p", 3000, "Port")
	flag.StringVar(&configFile, "c", "config.sample.ini", "AppConfig File")
	flag.Parse()
	hostAddress = fmt.Sprintf("%s:%d", host, port)
}

func createApp(config *config.HoaxConfig) error {
	var err error

	// global logger setup
	err = logger.CreateLogger()
	if err != nil {
		fmt.Println("logger init error")
		return err
	}

	// http server setup
	err = http.CreateHTTPServer()
	if err != nil {
		fmt.Println("HTTP init error")
		return err
	}
	routes.InitRoutes()

	// database setup
	if config.Db != "" {
		err = database.InitDatabase()
	}
	if err != nil {
		fmt.Println("DB init error")
		return err
	}
	return nil
}

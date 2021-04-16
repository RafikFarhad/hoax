package main

import (
	"flag"
	"fmt"
	"github.com/RafikFarhad/hoax/app"
	"github.com/RafikFarhad/hoax/config"
	_ "github.com/RafikFarhad/hoax/docs"
	"github.com/RafikFarhad/hoax/http"
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

	// create app
	if err := app.CreateApp(true); err != nil {
		panic(err)
	}

	// start Http Server
	err := http.AppHttp.Listen(config.AppConfig.Address)
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

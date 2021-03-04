package main

import (
	"flag"
	"fmt"
	"github.com/RafikFarhad/hoax/app"
	"github.com/RafikFarhad/hoax/config"
)

var (
	configFile  string
	hostAddress string
)

func main() {
	// Parse cmd args
	parseInputArgs()

	// Parse .env
	hoaxConfig, err := config.ParseConfig(hostAddress, configFile)
	if err != nil {
		panic(fmt.Sprintf("[INI] %+v", err))
	}
	// Create app
	if err := app.CreateApp(hoaxConfig); err != nil {
		panic(err)
	}

	// Start Http Server
	app.App.Http.Listen(hoaxConfig.Address)
}

func parseInputArgs() {
	var host string
	var port int
	// Parse cmd args
	flag.StringVar(&host, "h", "127.0.0.1", "Host")
	flag.IntVar(&port, "p", 3000, "Port")
	flag.StringVar(&configFile, "c", "config.sample.ini", "Config File")
	flag.Parse()
	hostAddress = fmt.Sprintf("%s:%d", host, port)
}

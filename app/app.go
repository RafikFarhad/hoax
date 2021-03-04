package app

import (
	"fmt"
	"github.com/RafikFarhad/hoax/config"
	"github.com/RafikFarhad/hoax/database"
	"github.com/RafikFarhad/hoax/http"
	"github.com/RafikFarhad/hoax/routes"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Hoax struct {
	Http   *fiber.App
	Config *config.HoaxConfig
	Db     *gorm.DB
}

var App *Hoax

func CreateApp(config *config.HoaxConfig) error {
	var err error
	App = &Hoax{Config: config}

	// Http server setup
	App.Http, err = http.CreateHTTPServer(config)
	if err != nil {
		fmt.Println("HTTP init error")
		return err
	}
	routes.InitRoutes(App.Http, App.Config)

	// Database setup
	if config.DbHost != "" {
		App.Db, err = database.InitMySqlDb(config)
	}
	if err != nil {
		fmt.Println("DB init error")
		return err
	}
	return nil
}

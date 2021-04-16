package app

import (
	"errors"
	"fmt"
	"github.com/RafikFarhad/hoax/cache"
	"github.com/RafikFarhad/hoax/database"
	"github.com/RafikFarhad/hoax/database/model"
	"github.com/RafikFarhad/hoax/http"
	"github.com/RafikFarhad/hoax/logger"
	"github.com/RafikFarhad/hoax/routes"
)

func CreateApp(autoMigration bool) error {
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

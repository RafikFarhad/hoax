package app

import (
	"github.com/RafikFarhad/hoax/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Hoax struct {
	Http   *fiber.App
	Config *config.HoaxConfig
	Db     *gorm.DB
}

var App *Hoax

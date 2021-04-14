package logger

import (
	"github.com/RafikFarhad/hoax/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func InitLog() error {
	logConfig := config.AppConfig.LogConfig
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	level, _ := zerolog.ParseLevel(logConfig.Level)
	zerolog.SetGlobalLevel(level)
	switch logConfig.Agent {
	case "file":
		logFile, err := os.OpenFile(logConfig.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		log.Logger = log.Output(logFile)
		break
	case "std":
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		break
	default:
		log.Logger = zerolog.Nop()
	}
	return nil
}

package logger

import (
	"log"
	"os"
)

var GlobalLogger *log.Logger

func CreateLogger() error {
	GlobalLogger = log.New(os.Stdout, "hoax :: ", log.LstdFlags)
	return nil
}

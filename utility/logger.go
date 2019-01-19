package utility

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func Logger(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

		// TODO: Configurable logger settings
		// Setting up Logger
		logPath := "hoax.log"
		logWriter := logWriter(logPath)
		logger := log.New(
			logWriter,
			"hoax:: ",
			log.Ldate|log.Ltime|log.Lshortfile,
		)
		// Logging this request
		logger.Printf("%s %s %s\n", req.URL, req.Method, req.RemoteAddr)
		// Putting logger in the request context so that it can be
		// accessed in the request life cycle
		ctx := context.WithValue(req.Context(), "logger", logger)
		req = req.WithContext(ctx)
		next(w, req, params)
	}
}

func logWriter(logPath string) *os.File {
	if logPath == "" {
		logPath = "hoax.log"
	}

	logFile, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

	if err != nil {
		log.Fatal("Can not open LOG FILE:", err)
	}
	return logFile
}

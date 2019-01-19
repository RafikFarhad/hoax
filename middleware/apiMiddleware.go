package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func ApiMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		now := time.Now()
		next(w, req, params)
		// TODO:: Logger Example, To be removed in future
		if logger := req.Context().Value("logger"); logger != nil {
			logger.(*log.Logger).Printf(
				"API:: Response Time: %dns",
				time.Since(now).Nanoseconds())
		}
	}
}

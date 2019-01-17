package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HomePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome to the world of speed!")
	fmt.Fprintf(w, "\n")
	// fmt.Fprintf(w, req.URL.Path)
}

func WebRoutes(router *httprouter.Router) {
	router.GET("/", HomePage)
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RafikFarhad/hoax/middleware"
	"github.com/RafikFarhad/hoax/router"
	"github.com/RafikFarhad/hoax/utility"
	"github.com/julienschmidt/httprouter"
)

func main() {

	fmt.Println("Serving GO...")

	mainRouter := httprouter.New()

	for route, handler := range router.WebRoutes() {
		mainRouter.GET("/"+route, utility.Logger(handler))
	}

	for route, handler := range router.ApiRoutes() {
		mainRouter.GET(
			"/api/"+route,
			utility.Logger(
				middleware.ApiMiddleware(handler)))

		fmt.Println("/api/" + route)
	}

	log.Fatal(http.ListenAndServe(":5000", mainRouter))
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RafikFarhad/hoax/router"
	"github.com/julienschmidt/httprouter"
)

func apiHandler() *http.ServeMux {
	apiServeMux := http.NewServeMux()
	apiServeMux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the API home page!")
	})
	apiServeMux.HandleFunc("/api/a", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the API/A home page!")
	})
	return apiServeMux
}

func HomePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome to the world of speed!")
	fmt.Fprintf(w, "\n")
	// fmt.Fprintf(w, req.URL.Path)
}

func main() {

	fmt.Println("Serving GO...")

	mainRouter := httprouter.New()

	router.WebRoutes(mainRouter)

	for route, function := range router.ApiRoutes() {
		mainRouter.GET("/api/"+route, function)
		fmt.Println("/api/" + route)
	}

	log.Fatal(http.ListenAndServe(":5000", mainRouter))
}

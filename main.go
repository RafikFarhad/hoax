package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RafikFarhad/hoax/router"
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

func bpiHandler() *http.ServeMux {
	apiServeMux := http.NewServeMux()
	apiServeMux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the BPI home page!")
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, req.URL.Path)
	})
	apiServeMux.HandleFunc("/bpi/1", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the BPI/1 home page!")
	})
	return apiServeMux
}

func main() {

	fmt.Println("Serving GO...")

	serveMux := http.NewServeMux()

	serveMux.Handle("/", router.GetWebRouter())
	serveMux.Handle("/api/", router.GetApiRouter())

	// serveMux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	// The "/" pattern matches everything, so we need to check
	// 	// that we're at the root here.
	// 	if req.URL.Path != "/" {
	// 		http.NotFound(w, req)
	// 		return
	// 	}
	// 	fmt.Fprintf(w, "Welcome to the home page!")
	// })
	serveMux.HandleFunc("/p", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "pppppppp")
	})
	serveMux.HandleFunc("/q", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "qqqqqqqqqqqq")
	})
	log.Fatal(http.ListenAndServe(":5000", serveMux))
}

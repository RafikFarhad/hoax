package router

import (
	"fmt"
	"net/http"
)

var webRouter *http.ServeMux

func GetWebRouter() *http.ServeMux {
	webRouter = http.NewServeMux()
	webRouter.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; UTF-8")
		fmt.Fprintf(w, "Welcome to the home page!\n")
		fmt.Fprintf(w, "Your Route is: \""+req.URL.Path+"\"")
	})
	return webRouter
}

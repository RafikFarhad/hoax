package router

import (
	"encoding/json"
	"net/http"

	"github.com/RafikFarhad/hoax/controller/books"
	"github.com/RafikFarhad/hoax/response"
)

var apiRouter *http.ServeMux

func GetApiRouter() *http.ServeMux {

	apiRouter = http.NewServeMux()

	apiRouter.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := response.ApiResponse{}
		response.Message = "API is working. " + "Your Route is: \"" + req.URL.Path + "\""
		response.Status = true
		json.NewEncoder(w).Encode(response)
	})

	apiRouter.HandleFunc("/api/books", books.Index)
	return apiRouter
}

package router

import (
	"encoding/json"
	"net/http"

	"github.com/RafikFarhad/hoax/controller/books"

	"github.com/RafikFarhad/hoax/response"
	"github.com/julienschmidt/httprouter"
)

type RoutesMapper map[string]func(w http.ResponseWriter, req *http.Request, _ httprouter.Params)

func ping(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if req.URL.Path == "/api/" {
		http.Redirect(w, req, "ping", 301)
	}
	w.Header().Set("Content-Type", "application/json")
	response := response.ApiResponse{}
	// response.Message = "Pong Pong!"
	// response.Status = true
	response["message"] = "Pong Pong!"
	response["status"] = true
	json.NewEncoder(w).Encode(response)
}

func ApiRoutes() RoutesMapper {
	routes := make(RoutesMapper)
	routeList(routes)
	return routes
}

func routeList(routes RoutesMapper) {
	routes[""] = ping
	routes["ping"] = ping
	routes["books"] = books.Index
}

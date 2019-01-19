package api_index

import (
	"encoding/json"
	"net/http"

	"github.com/RafikFarhad/hoax/response"
	"github.com/julienschmidt/httprouter"
)

func Ping(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if req.URL.Path == "/api/" {
		http.Redirect(w, req, "ping", 301)
	}
	response := response.ApiResponse{}
	response["message"] = "Pong Pong!"
	response["status"] = true
	json.NewEncoder(w).Encode(response)
}

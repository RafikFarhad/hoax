package books

import (
	"encoding/json"
	"net/http"

	"github.com/RafikFarhad/hoax/response"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := response.ApiResponse{}
	response.Message = "books index"
	response.Status = true
	json.NewEncoder(w).Encode(response)
}

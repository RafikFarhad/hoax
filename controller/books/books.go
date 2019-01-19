package books

import (
	"encoding/json"
	"net/http"

	"github.com/RafikFarhad/hoax/model"
	"github.com/RafikFarhad/hoax/response"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := response.ApiResponse{}
	book := []model.Book{
		{Title: "Awesome Book", Id: 1, Isbn: "100/200", Price: 100},
		{Title: "Not A Awesome Book", Id: 2, Isbn: "a/b", Price: -100},
	}
	response["books"] = book
	json.NewEncoder(w).Encode(response)
}

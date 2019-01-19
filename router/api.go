package router

import (
	"github.com/RafikFarhad/hoax/controller/api_index"
	"github.com/RafikFarhad/hoax/controller/books"

	"github.com/julienschmidt/httprouter"
)

type RoutesMapper map[string]httprouter.Handle

func ApiRoutes() RoutesMapper {
	routes := make(RoutesMapper)
	apiRouteList(routes)
	return routes
}

func apiRouteList(routes RoutesMapper) {
	routes[""] = api_index.Ping
	routes["ping"] = api_index.Ping
	routes["books"] = books.Index
}

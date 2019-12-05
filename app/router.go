package app

import (
	"garen/logger"
	"net/http"
	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

type Route struct {
	Name	string
	Method	string
	Pattern	string
	HandlerFunc	http.HandlerFunc
}

type Routers []Route

var routers = Routers{
	Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: controller.Index,
	},
	Route{
		Name:        "AddAlbum",
		Method:      "POST",
		Pattern:     "/",
		HandlerFunc: controller.addAlbum,
	},
	Route{
		Name:        "UpdateAlbum",
		Method:      "PUT",
		Pattern:     "/",
		HandlerFunc: controller.UpdateAlbum,
	},
	Route{
		Name:        "DeleteAlbum",
		Method:      "DELETE",
		Pattern:     "/{id}",
		HandlerFunc: controller.DeleteAlbum,
	},
}

func NewRouter() *mux.Router  {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routers {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

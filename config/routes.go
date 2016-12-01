package config

import (
	"net/http"

	"github.com/KoseSoftware/secret-santa-api/controllers"
	"github.com/gorilla/mux"
)

type Route struct {
	Pattern     string
	HandlerFunc http.HandlerFunc
	Method      string
	Name        string
}

type Routes []Route

func NewRouter(bc *controllers.BaseController) *mux.Router {
	hc := controllers.NewHomepageController(bc)

	var routes = Routes{
		Route{"/", hc.Index, "GET", "homepage"},
		Route{"/test", hc.Test, "GET", "test"},
	}

	router := mux.NewRouter().StrictSlash(false)
	for _, route := range routes {
		router.
			Path(route.Pattern).
			Handler(route.HandlerFunc).
			Methods(route.Method).
			Name(route.Name)
	}

	return router
}

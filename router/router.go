package router

import (
	"net/http"

	controllers "github.com/backend-services/controllers"
	mux "github.com/gorilla/mux"
)

// New returns a new router
func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		if route.Protected {
			router.Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HandlerFunc)
		} else {
			router.Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HandlerFunc)
		}
	}
	return router
}

// NewHealthcheckRouter ... New health check router
func NewHealthcheckRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").
		Path("/healthstatus").
		Name("Healthcheck api").
		Handler(http.HandlerFunc(controllers.HealthCheck))
	return router
}

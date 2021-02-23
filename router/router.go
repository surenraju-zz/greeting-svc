package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/greeting-svc/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = LogRequest(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	router.Use(middleware.PrometheusMiddleware)
	router.Path("/metrics").Handler(promhttp.Handler())
	return router
}

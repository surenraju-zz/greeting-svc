package router

import (
	"net/http"

	"github.com/greeting-svc/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"Greet",
		"GET",
		"/api/v1/greet/{name}",
		handler.Greet,
	},
}

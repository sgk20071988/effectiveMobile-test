package router

import (
	"effectiveMobileTest/internal/handlers"
	"net/http"
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
		"GetCars",
		"GET",
		"/cars",
		handlers.GetCars,
	},
	Route{
		"InsertCar",
		"POST",
		"/newcar",
		handlers.InsertCar,
	},
}

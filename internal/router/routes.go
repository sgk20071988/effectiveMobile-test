package router

import (
	"effectiveMobileTest/internal/handlers"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
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
		"/cars/",
		handlers.GetCars,
	},
	Route{
		"GetCar",
		"GET",
		"/cars/{regNum}",
		handlers.GetCar,
	},
	Route{
		"InsertCar",
		"POST",
		"/cars/",
		handlers.InsertCar,
	},
	Route{
		"UpdateCar",
		"PATCH",
		"/cars/{regNum}",
		handlers.UpdateCar,
	},
	Route{
		"DeleteCar",
		"DELETE",
		"/cars/{regNum}",
		handlers.DeleteCar,
	},
	Route{
		"swagger",
		"GET",
		"/swagger/",
		httpSwagger.Handler(httpSwagger.URL("http://localhost:1223/swagger/doc.json")),
	},
}

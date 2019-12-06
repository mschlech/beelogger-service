package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Ïndex",
		"GET",
		"/Weight/{hiveId}",
		getWeight,
	},
	Route{
		"Ïndex",
		"GET",
		"/temperature/{hiveId}",
		getTemperature,
	},
}

package main

import (
	"github.com/mschlech/beelogger-service/cmd/handler"
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
		"Weight",
		"GET",
		"/Weight/{hiveId}",
		handler.GetWeight,
	},
	Route{
		"Temperature",
		"GET",
		"/temperature/{hiveId}",
		handler.GetTemperature,
	},
	Route{
		Name:        "Humidity",
		Method:      "GET",
		Pattern:     "/humidity/{hiveId}",
		HandlerFunc: handler.GetHumidity,
	},
	Route{
		Name:        "Gas",
		Method:      "GET",
		Pattern:     "/gas/{hiveId}",
		HandlerFunc: handler.GetGasIndication,
	},
	Route{
		Name:        "Lux",
		Method:      "GET",
		Pattern:     "/lux/{hiveId}",
		HandlerFunc: handler.GetLux,
	},

	Route{
		Name:        "BatteryConsumption",
		Method:      "GET",
		Pattern:     "/battery/{hiveId}/consumption",
		HandlerFunc: handler.GetBatteryConsumption,
	},
	Route{
		Name:        "GetHiveStatus",
		Method:      "GET",
		Pattern:     "/hives/{hiveId}/status",
		HandlerFunc: handler.GetHiveStatus,
	},
	Route{
		Name:        "GetHiveIoTStatus",
		Method:      "GET",
		Pattern:     "/hives/{hiveId}/status",
		HandlerFunc: handler.GetHiveIoTStatus,
	},


}

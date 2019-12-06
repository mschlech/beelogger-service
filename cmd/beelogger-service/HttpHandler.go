package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getWeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var hiveId int
	var err error
	log.Print(err, hiveId, vars)

	return
}

func getTemperature(w http.ResponseWriter, r *http.Request) {

	return
}

//get the metrics from the iot device via http (@TODO switch to MQTT)

func postTemperature(w http.ResponseWriter, r *http.Request) {

	return
}

func postWeight(w http.ResponseWriter, r *http.Request) {

	return
}

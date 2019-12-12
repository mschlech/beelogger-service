package handler

import (
	"github.com/gorilla/mux"
	"github.com/mschlech/beelogger-service/cmd/models"
	"log"
	"net/http"
)

var hiveId int

func GetWeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	log.Print(err, hiveId, vars)

	return
}

func GetTemperature(w http.ResponseWriter, r *http.Request) {

	return
}

func GetHumidity(w http.ResponseWriter, r *http.Request) {

	return
}

func GetGasIndication(w http.ResponseWriter, r *http.Request) {

	return
}

// json payload responses with all metrics
func GetHiveStatus( w http.ResponseWriter, r *http.Request) {
	var hive []models.Hive
	jsonResponse(w,http.StatusOK,hive)
}

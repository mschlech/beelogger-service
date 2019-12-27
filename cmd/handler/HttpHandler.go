package handler

import (
	"github.com/gorilla/mux"
	"github.com/mschlech/beelogger-service/cmd/dao"
	"github.com/mschlech/beelogger-service/cmd/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

var hiveId int

var hive models.Hive
var hiveDao=dao.HiveDao{}

func init() {
	hiveDao.Server = "127.0.0.1:27017"
	hiveDao.Database = "hive"
	log.Print("Init executed")
	hiveDao.Connect()
}

func GetWeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	log.Print(err, hiveId, vars)

	return
}

func GetTemperature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	log.Print(
		id)


	return
}

func GetHumidity(w http.ResponseWriter, r *http.Request) {

	return
}

func GetGasIndication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var (
		hive models.Hive
	)

	log.Print(params)
	jsonResponse(w, 200, hive)
}

// json payload responses with all metrics
func GetHiveStatus( w http.ResponseWriter, r *http.Request) {
	var hive []models.Hive
	jsonResponse(w,http.StatusOK,hive)
}

func GetAllHiveMetrics(w http.ResponseWriter , r *http.Request) {
	log.Print("getAllHiveMetrics " , w)
	hives, error := hiveDao.GetAllHivesMetrics()
	log.Print("sample : ", hives,error)

	if error != nil {
		log.Fatal("error", error)

	}
	log.Print("from HttpHandler hives -> " , hives)
	jsonResponse(w, http.StatusOK, hives )
}


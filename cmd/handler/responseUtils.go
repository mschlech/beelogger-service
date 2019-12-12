package handler

import (
	"encoding/json"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	response,error := json.Marshal(payload)
	if error!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

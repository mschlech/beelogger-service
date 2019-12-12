package beelogger

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("start beelogger service")
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}

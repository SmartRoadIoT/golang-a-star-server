package main

import (
	"log"
	"net/http"
	"smart-road/main/model/controller"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/distance", controller.HandleDistanceRequest)
	mux.HandleFunc("/route/vehicle", controller.HandleFindRouteRequest)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
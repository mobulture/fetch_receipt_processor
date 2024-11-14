package main

import (
	"fetch_receipt_processor/src/cache"
	"fetch_receipt_processor/src/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRoutes()
	cache.InitCache()
	log.Println("Starting server")
	port := ":8080"
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Println(port, "is already in use")
	}
}

package main

import (
"net/http"
"log"
"fetch_receipt_processor/src/routes"
"fetch_receipt_processor/src/cache"
)


func main() {
    router := routes.SetupRoutes()
	cache.InitCache()
	log.Println("Staring server")
	http.ListenAndServe(":8080",router)
}


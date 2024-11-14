package main

import (
"net/http"
"log"
"fetch_receipt_processor/src/routes"
)


func main() {
    router := routes.SetupRoutes()
	log.Println("Staring server")
	http.ListenAndServe(":8080",router)
}


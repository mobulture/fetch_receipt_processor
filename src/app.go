package main

import (
"net/http"
"log"

)

func main() {
    router := SetupRoutes()
	log.Println("Staring server")
	http.ListenAndServe(":8080",router)
}


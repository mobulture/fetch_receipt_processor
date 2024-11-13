package main

import (
 "encoding/json"
 "net/http"
 "log"
 "github.com/google/uuid"
 
)


func PostReceipt(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type", "application/json")
 log.Println("In post receipt")
 response := map[string]interface{}{
	"ID": 2024,
}
 json.NewEncoder(w).Encode(response)
}

func GetPoints(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	log.Println("in get points")
	response := map[string]interface{}{
		"Points": 2023,
	}
	json.NewEncoder(w).Encode(response)

}
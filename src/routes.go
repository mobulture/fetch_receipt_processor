package main
import (
	"github.com/gorilla/mux" 
)
   func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
   
	router.HandleFunc("/receipts/process", PostReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", GetPoints).Methods("GET")

	return router
   }
   
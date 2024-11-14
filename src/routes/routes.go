package routes
import (
	"fetch_receipt_processor/src/handlers"
	"github.com/gorilla/mux" 
)
   func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
   
	router.HandleFunc("/receipts/process", handlers.PostReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")

	return router
   }
   
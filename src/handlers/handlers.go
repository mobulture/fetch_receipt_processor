package handlers

import (
	"bytes"
	"encoding/json"
	"fetch_receipt_processor/src/cache"
	"fetch_receipt_processor/src/types"
	"fetch_receipt_processor/src/utils"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func PostReceipt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Received postReceipt request")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	var receipt types.Receipt
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&receipt)
	UUID, err := utils.GenerateUUID(receipt) // if we create the UUID before points are calculated it should be fine
	if err != nil {
		http.Error(w, "Could not generate ID", http.StatusBadRequest)
		errorResponse := map[string]string{
			"error":  "Could not generate ID",
			"status": "400",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	response := map[string]interface{}{
		"ID": UUID,
	}
	json.NewEncoder(w).Encode(response)
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	receipt, exists := cache.Get(id)
	if exists == false {
		http.Error(w, "ID not in cache ", http.StatusBadRequest)
		errorResponse := map[string]string{
			"error":  "ID not found in cache",
			"status": "400",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	log.Println("Found corresponding receipt to id", id)
	response := map[string]interface{}{
		"Points": receipt.Points,
	}
	json.NewEncoder(w).Encode(response)
}

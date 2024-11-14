package handlers

import (
 "encoding/json"
 "net/http"
 "log"
 "io"
 "bytes"
 "github.com/google/uuid"
 "fetch_receipt_processor/src/types"
)



func generateUUID(receipt types.Receipt) (string, error) {
	data, err := json.Marshal(receipt)
	namespace := uuid.NameSpaceURL
	newUUID := uuid.NewSHA1(namespace, data)
	log.Println(newUUID)
	return newUUID.String(), err
}

func PostReceipt(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type", "application/json")
 log.Println("In post receipt")

 body, err := io.ReadAll(r.Body)
 if err != nil {
	 http.Error(w, "Unable to read request body", http.StatusInternalServerError)
	 return
 }

 log.Println("Received request body:", string(body))
 
 r.Body = io.NopCloser(bytes.NewBuffer(body))

 var receipt types.Receipt
 decoder := json.NewDecoder(r.Body)
 err = decoder.Decode(&receipt)
 var UUID string
 UUID,err = generateUUID(receipt)
 if err != nil {
	http.Error(w, "Invalid ", http.StatusBadRequest)
	return
}
 response := map[string]interface{}{
	"ID": UUID,
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
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"io"
	"fmt"
	"fetch_receipt_processor/src/types"
	"testing"
	"github.com/stretchr/testify/assert"
)

type PostResponse struct {
	ID string 
}

type GetResponse struct {
	Points int
}


func TestUUIDGen (t *testing.T){
	// Tests that UUID generation is unique

	receipt1 := types.Receipt{
		Retailer:    "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: [] types.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: "12.00"},
		},
		Total: "35.35",
	}
	receipt2 := types.Receipt{
		Retailer:    "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []types.Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
		Total: "9.00",
	}
	// Generates UUID of first receipt
	jsonData, err:= json.Marshal(receipt1)

	resp, err := http.Post("http://127.0.0.1:8080/receipts/process", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()
	
	var postResponse PostResponse
	respBody, err := io.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(string(respBody)), &postResponse)
	log.Println("ID:", postResponse.ID)

	firstID:= postResponse.ID

	// Generates UUID of second receipt
	jsonData, err= json.Marshal(receipt2)
	resp, err = http.Post("http://127.0.0.1:8080/receipts/process", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()
	
	respBody, err = io.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(string(respBody)), &postResponse)
	log.Println("ID:", postResponse.ID)
	secondID := postResponse.ID

	// Asserts that they are both not equal to test unique generation
	assert.NotEqual(t, firstID, secondID, "Generated UUIDs should not be equal")

}

func TestPointsFirst(t *testing.T){
	//Based off first testcase, should also test caching functionality as this receipt is passed in earlier test
	receipt := types.Receipt{
		Retailer:    "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: [] types.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: "12.00"},
		},
		Total: "35.35",
	}
	jsonData, err:= json.Marshal(receipt)

	resp, err := http.Post("http://127.0.0.1:8080/receipts/process", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	var postResponse PostResponse
	respBody, err := io.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(string(respBody)), &postResponse)
	log.Println("ID:", postResponse.ID)

	url := fmt.Sprintf("http://127.0.0.1:8080/receipts/%s/points", postResponse.ID)


	var getResponse GetResponse
	resp, err = http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	respBody, err = io.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(string(respBody)), &getResponse)
	log.Println("Assert points:", getResponse.Points ,"= 28")
	assert.Equal(t, 28, getResponse.Points)


}
func TestPointsSecond(t *testing.T) {
	//Based off second testcase, should also test caching functionality as this receipt is passed in earlier test
	receipt := types.Receipt{
		Retailer:    "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []types.Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
		Total: "9.00",
	}
	jsonData, err:= json.Marshal(receipt)

	resp, err := http.Post("http://127.0.0.1:8080/receipts/process", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	var postResponse PostResponse
	respBody, err := io.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(string(respBody)), &postResponse)
	log.Println("ID:", postResponse.ID)

	url := fmt.Sprintf("http://127.0.0.1:8080/receipts/%s/points", postResponse.ID)


	var getResponse GetResponse
	resp, err = http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	respBody, err = io.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(string(respBody)), &getResponse)
	log.Println("Assert points:", getResponse.Points ,"= 109")
	assert.Equal(t, 109, getResponse.Points)

}

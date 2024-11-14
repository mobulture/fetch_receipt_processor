package utils

import (
	"encoding/json"
	"log"
	"github.com/google/uuid"
	"fetch_receipt_processor/src/types"
	"fetch_receipt_processor/src/cache"
	"unicode"
	"strconv"
	"time"
	"math"
   )
   
func GenerateUUID(receipt types.Receipt) (string, error) {
	data, err := json.Marshal(receipt)
	namespace := uuid.NameSpaceURL
	UUID := uuid.NewSHA1(namespace, data).String()
	_,exists := cache.Get(UUID)
	if exists == true {
		log.Println("Receipt already has existing UUID, returning previously created UUID")
	} else {
		// If the ruleset for points does not change, it should be constant, so I wanted to preprocess it when a receipt is put in our cache
	 	//our UUID should be unique regardless of when we calculate the points
		receipt.Points= calculatePoints(receipt)
		log.Println("Generated new UUID: ",	UUID,", with ", receipt.Points," points")
		cache.Set(UUID,receipt)
	}
	return UUID, err
}

func countAlphaNumeric(s string) int{
	count:=0
	for _,r:= range s{
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			count++
		}
	}
	return count
}

func pointsFromItems(items []types.Item) int{
	points:= 0
	points+= (len(items) / 2) * 5 
	for _,item := range items{
		itemDescription:= item.ShortDescription
		itemPrice := item.Price

		if len(itemDescription) % 3 == 0 {
			itemPoints,_:= strconv.ParseFloat(itemPrice,64)
			itemPoints *= .2 
			points+= int(math.Ceil(itemPoints))
		}
	}

	return points
}

func pointsFromDate(date string) int{
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Println("Error parsing date:", err)
		return 0 
	}
	if parsedDate.Day() % 2 ==1 {
		return 6 
	}else {
		return 0 
	}

}

func pointsFromTime(timeStr string) int {
	parsedTime,err := time.Parse("15:04", timeStr)
	if err != nil {
		log.Println("Error parsing time:", err)
		return 0 
	}
	start,_ := time.Parse("15:04","14:00")
	end,_ := time.Parse("15:04","16:00")

	if parsedTime.After(start) && parsedTime.Before(end){
		return 10
	} else{
		return 0 
	}
}

func pointsFromTotal(totalString string) int {
	total, err := strconv.ParseFloat(totalString, 64)
    if err != nil {
        log.Println("Error converting string to float:", err)
	}
	points := 0

	if float64(int(total)) == total{
		points+=50
	}
	if math.Mod(total,0.25) == 0{
		points+=25
	}
	return points
}

func calculatePoints(receipt types.Receipt) (int){
	retailer:= receipt.Retailer
	dateString:= receipt.PurchaseDate
	timeString:= receipt.PurchaseTime
	items:= receipt.Items
	total:= receipt.Total

	points := 0
	points+= countAlphaNumeric(retailer)
	points+= pointsFromDate(dateString)
	points+= pointsFromTotal(total)
	points+= pointsFromTime(timeString)
	points+= pointsFromItems(items)
	return points
}

# How to Run


I decided to both code this in Go and decided to Dockerize the application. You can run the application in both ways

## Go

From the root of the directory, you can run `cd src`, then `go build`, then `go run .`. This should host an api on localhost:8080, you can send curl requests in the specified in the assessment requirements

An example is ` curl -X POST http://127.0.0.1:8080/receipts/process -d '{ "retailer": "Target", "purchaseDate": "2022-01-01", "purchaseTime": "13:01", "items": [ { "shortDescription": "Mountain Dew 12PK", "price": "6.49" },{ "shortDescription": "Emils Cheese Pizza", "price": "12.25" },{ "shortDescription": "Knorr Creamy Chicken", "price": "1.26" },{ "shortDescription": "Doritos Nacho Cheese", "price": "3.35" },{ "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ", "price": "12.00" } ], "total": "35.35" }' -H "Content-Type: application/json"`

## Docker

To run the application with docker, you can simply run `docker-compose up --build` in the root directory, which will then build the docker container and host the application upon the completion of the build.


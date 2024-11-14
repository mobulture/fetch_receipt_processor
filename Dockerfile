FROM golang:1.23
 
WORKDIR /app
 
COPY . .
 
RUN go build -o fetch_receipt_processor ./src

EXPOSE 8080
 
CMD ["./fetch_receipt_processor"]
FROM golang:1.21

# Expose the port that the application listens on.
WORKDIR /src

COPY . . 

EXPOSE 8080

RUN go build . 

CMD ["./src"]

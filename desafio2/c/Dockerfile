FROM golang:1.15

WORKDIR /go/src/microservice-c

COPY . .

RUN go mod download

RUN GOOS=linux go build c.go

EXPOSE 9092

ENTRYPOINT ["./c"]
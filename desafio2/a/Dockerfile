FROM golang:1.15

WORKDIR /go/src/microservice-a

COPY . .

RUN go mod download

RUN GOOS=linux go build a.go

EXPOSE 9090

ENTRYPOINT ["./a"]

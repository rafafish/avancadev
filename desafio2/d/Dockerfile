FROM golang:1.15

WORKDIR /go/src/microservice-d

COPY . .

RUN go mod download

RUN GOOS=linux go build d.go

EXPOSE 9093

ENTRYPOINT ["./d"]
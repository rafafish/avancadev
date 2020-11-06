package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"github.com/wesleywillians/go-rabbitmq/queue"
)

type Result struct {
	Status string
}

type Order struct {
	ID       uuid.UUID
	Coupon   string
	CcNumber string
}

func NewOrder() Order {
	return Order{ID: uuid.NewV4()}
}

const (
	InvalidCoupon   = "invalid"
	ValidCoupon     = "valid"
	ConnectionError = "connection error"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}
}

func main() {
	messageChannel := make(chan amqp.Delivery)
	rabittMQ := queue.NewRabbitMQ()
	ch := rabittMQ.Connect()
	defer ch.Close()

	rabittMQ.Consume(messageChannel)

	for msg := range messageChannel {
		process(msg)
	}
}

func process(msg amqp.Delivery) {

	order := NewOrder()

	json.Unmarshal(msg.Body, &order)

	resultCoupon := makeHttpCall("http://localhost:9092", order.Coupon)

	switch resultCoupon.Status {
	case InvalidCoupon:
		log.Println("Order: ", order.ID, ": Invalid Coupon")
	case ConnectionError:
		msg.Reject(false)
		log.Println("Order: ", order.ID, ": Could Not Process")
	case ValidCoupon:
		log.Println("Order: ", order.ID, ": Processed")
	}

}

func makeHttpCall(urlMicroservice string, coupon string) Result {

	values := url.Values{}
	values.Add("coupon", coupon)

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 0

	res, err := retryClient.PostForm(urlMicroservice, values)
	if err != nil {
		result := Result{Status: ConnectionError}
		return result
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error processing result")
	}

	result := Result{}

	json.Unmarshal(data, &result)

	return result

}

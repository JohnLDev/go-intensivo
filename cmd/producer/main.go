package main

import (
	"encoding/json"
	"math/rand"

	"github.com/google/uuid"
	"github.com/johnldev/imersao-golang/src/order/entity"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(ch *amqp.Channel, order entity.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"amq.direct", // exchange default
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return err
	}
	return nil
}

func GenerateOrders() entity.Order {
	return entity.Order{
		ID:    uuid.New().String(),
		Tax:   rand.Float64() * 10,
		Price: rand.Float64() * 100,
	}
}
func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	defer channel.Close()

	for i := 0; i < 10; i++ {
		Publish(channel, GenerateOrders())
	}
}

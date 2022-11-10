package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/johnldev/imersao-golang/pkg/rabbitmq"
	"github.com/johnldev/imersao-golang/src/order/infra/database"
	"github.com/johnldev/imersao-golang/src/order/useCases"
	"github.com/rabbitmq/amqp091-go"

	//sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepoistory(db)
	useCase := useCases.CalculateFinalPriceUseCase{
		OrderRepository: repository,
	}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp091.Delivery) // * channel
	go rabbitmq.Consume(ch, out)       // * T2

	qtdWorkers := 100
	for i := 1; i <= qtdWorkers; i++ {
		go worker(out, useCase, i) // criando mais threads
	}
	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		getTotalUseCase := useCases.NewGetTotalUseCase(repository)
		output, err := getTotalUseCase.Execute()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		json.NewEncoder(w).Encode(output)
	})
	http.ListenAndServe(":8080", nil) //Cria uma nova thread por requisição
}

func worker(deliveryMessage <-chan amqp091.Delivery, useCase useCases.CalculateFinalPriceUseCase, workerId int) {
	for msg := range deliveryMessage {
		var input useCases.CalculateOrderInputDTO
		err := json.Unmarshal(msg.Body, &input)

		if err != nil {
			fmt.Println(err.Error())

			panic(err)
		}
		output, err := useCase.Execute(input)
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
		err = msg.Ack(false)
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
		fmt.Printf("Worker %d has processed the order %s \n", workerId, output.ID)
		time.Sleep(1 * time.Second)
		// fmt.Println(output)
	}
}

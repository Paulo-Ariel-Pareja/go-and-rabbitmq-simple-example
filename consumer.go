package main

import (
	"fmt"

	"github.com/Paulo-Ariel-Pareja/go-and-rabbitmq-simple-example/helpers"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer iniciado")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	helpers.HandleErr(err)

	defer conn.Close()

	fmt.Println("Conectado a Rabbitmq")

	ch, err := conn.Channel()
	helpers.HandleErr(err)

	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recibiendo mensaje: %s\n", d.Body)
		}
	}()

	fmt.Println("Consumer conectado correctamente - esperando mensajes")
	<-forever
}

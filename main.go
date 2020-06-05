package main

import (
	"fmt"

	"github.com/Paulo-Ariel-Pareja/go-and-rabbitmq-simple-example/helpers"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Ejemplo con Go y Rabbitmq")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	helpers.HandleErr(err)

	defer conn.Close()

	fmt.Println("Conectado a Rabbitmq")

	ch, err := conn.Channel()
	helpers.HandleErr(err)

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	helpers.HandleErr(err)
	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World!"),
		},
	)
	helpers.HandleErr(err)

	fmt.Println("Mensaje publicado en la queue")
}

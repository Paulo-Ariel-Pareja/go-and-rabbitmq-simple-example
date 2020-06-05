package main

import (
	"fmt"

	"github.com/stradway/amqp"
)

func main() {
	fmt.Println("Ejemplo con Go y Rabbitmq")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nul {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()
}

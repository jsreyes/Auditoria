package main

import (
	"fmt"
	"log"
	//"strings"
	//"strconv"
	"github.com/streadway/amqp"
)

//Variables para la conexión con RabbitMQ y el canal correspondiente
var connection *amqp.Connection
var chl *amqp.Channel

func main() {
	//Dequeuing message actions
	msgs := ConsumeQueue("auditoria2")
	go func() {
		forever := make(chan bool)
		go func() {
			for d := range msgs {
				processMessage(d)
			}
		}()
		log.Printf(" [*] Waiting for messages actions.")
		<-forever
	}()

	running := make(chan bool)
	log.Printf("Para salir oprima Crtl C")
	<-running

}

func consumirCola(nombreCola string) <-chan amqp.Delivery {

	var uri = "amqp://guest:guest@10.20.0.175:5672/"
	//var channel =  "auditoria"

	con, err := amqp.Dial(uri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer con.Close()

	connection = con

	chanel, err := con.Channel()
	failOnError(err, "Failed to open a channel")
	defer chanel.Close()

	chl = chanel

	q, err := chl.QueueDeclare(
		nombreCola,
		true,
		false,
		false,
		false,
		nil,
	)

	log.Print(q)
	failOnError(err, "Failed to declare a queue")

	msgs, err := chl.Consume(
		nombreCola,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to register a consumer")

	return msgs

}

func ConsumeQueue(queuename string) <-chan amqp.Delivery {
	var uri = "amqp://guest:guest@10.20.0.175:5672/"
	//var channel =  "auditoria"

	con, err := amqp.Dial(uri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer con.Close()

	connection = con

	chanel, err := con.Channel()
	failOnError(err, "Failed to open a channel")
	defer chanel.Close()

	chl = chanel

	q, err := chl.QueueDeclare(
		queuename, // name
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	failOnError(err, "Failed to declare a queue")

	msgs, err := chl.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	return msgs
}

func processMessage(a amqp.Delivery) bool {
	fmt.Print(a)
	return true
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panic("%s: %s", msg, err)
		log.Panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

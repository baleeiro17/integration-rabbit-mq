package consumer

import (
	"github.com/streadway/amqp"
	"integration-rabbit-mq/internal/json"
	"integration-rabbit-mq/internal/logs"
	"log"
)

func Receive() {

	conn, err := amqp.Dial("amqp://guest:guest@10.254.2.61:5672/")
	logs.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	logs.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	logs.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	logs.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			p, err := json.Deserialize(d.Body)
			if err != nil {
				log.Printf("Error in read the message in client")
			} else {
				log.Printf("Received a message:")
				log.Printf("ID is: %s\n", p.Result[0].Id)
				log.Printf("Status is: %s\n", p.Result[0].Status)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

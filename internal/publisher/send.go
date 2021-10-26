package publisher

import (
	"github.com/streadway/amqp"
	"integration-rabbit-mq/internal/json"
	"integration-rabbit-mq/internal/logs"
	"integration-rabbit-mq/internal/repository"
)

// client for rabbit mq
func Send() {

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

	reserva, err := repository.GetReserva()
	logs.FailOnError(err, "Failed to get data in repository")

	info, err := json.Serialize(reserva)
	logs.FailOnError(err, "Failed to serialize the data")

	body := info
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	logs.FailOnError(err, "Failed to establish a message")

}

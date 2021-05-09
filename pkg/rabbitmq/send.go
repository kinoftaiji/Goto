package rabbitmq

import (
	"github.com/armnerd/go-skeleton/config"
	"log"

	"github.com/streadway/amqp"
)

func Send(topic string, body string) {
	// 获取参数
	host := config.RabbitMQConfig.Host
	port := config.RabbitMQConfig.Port
	user := config.RabbitMQConfig.User
	pass := config.RabbitMQConfig.Pass
	conn, err := amqp.Dial("amqp://" + user + ":" + pass + "@" + host + ":" + port + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"kungeek_direct", // name
		"direct",         // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	err = ch.Publish(
		"kungeek_direct", // exchange
		topic,            // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(body),
			DeliveryMode: 2,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

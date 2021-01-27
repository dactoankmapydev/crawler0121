package helper

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Rbmq struct {
	UserName string
	Password string
	Host     string
	Port     string
}

type Chan struct {
	Channel *amqp.Channel
}

func (rbmq *Rbmq) ConnectRbmq() (){
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", rbmq.UserName, rbmq.Password, rbmq.Host, rbmq.Port))
	if err != nil {
		log.Println("Failed to connect to RabbitMQ")
		log.Println(err)
	} else {
		fmt.Println("Success to connect to RabbitMQ")
	}

	ch, err := conn.Channel()
	failOnErr(err, "Failed to open a channel")
	defer ch.Close()
}

func (ch *Chan) Publish(routingKey string, data []byte) {
	//var ch *amqp.Channel
	q, err := ch.Channel.QueueDeclare(
		routingKey,
		true,
		false,
		false,
		false,
		nil,
		)
	fmt.Println(q)
	failOnErr(err, "Failed to create the queue")

	err = ch.Channel.Publish(
		"",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: data,
			Priority: 5,
		})
	failOnErr(err, "Failed to publish a msg")
	log.Printf("Successfully published message %s", data)
}

func failOnErr(err error, msg string) {
	if err != nil {
		fmt.Println("%s: %s", msg, err)
	}
}

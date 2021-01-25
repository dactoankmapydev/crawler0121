package helper

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Rbmq struct {
	Channel *amqp.Channel
	UserName string
	Password string
	Host     string
	Port     string
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

func (rbmq *Rbmq) Publish(routingKey string, data []byte) {
	err := rbmq.Channel.Publish(
		"",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: data,
		})
	failOnErr(err, "Failed to publish a msg")
	log.Printf("send %s", data)
}

func failOnErr(err error, msg string) {
	if err != nil {
		fmt.Println("%s: %s", msg, err)
	}
}

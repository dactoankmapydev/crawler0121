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

func (rbmq *Rbmq) ConnectRbmq() {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", rbmq.UserName, rbmq.Password, rbmq.Host, rbmq.Port))
	if err != nil {
		log.Println("Failed to connect to RabbitMQ")
		log.Println(err)
	} else {
		fmt.Println("Success to connect to RabbitMQ")
	}
	defer conn.Close()
}

package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func ConnectRbmq() (*amqp.Channel, error){
	rbmqHost := os.Getenv("RBMQ_HOST")
	rbmqPort := os.Getenv("RBMQ_PORT")
	rbmqUserName := os.Getenv("RBMQ_USER_NAME")
	rbmqPassword := os.Getenv("RBMQ_PASSWORD")
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", rbmqUserName, rbmqPassword, rbmqHost, rbmqPort))
	if err != nil {
		log.Println("Failed to connect to RabbitMQ")
		log.Println(err)
	} else {
		log.Println("Success to connect to RabbitMQ")
	}

	ch, err := conn.Channel()
	failOnErr(err, "Failed to open a channel")
	return ch, err
}

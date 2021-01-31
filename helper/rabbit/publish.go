package rabbit

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"ioc-provider/model"
	"log"
)

func PublishPost(routingKey string, data []model.Post) {
	ch, err := ConnectRbmq()
	q, err := ch.QueueDeclare(
		 routingKey,
		true,
		false,
		false,
		false,
		nil,
	)
	fmt.Println(q)
	failOnErr(err, "Failed to create the queue")

	for _, msg := range data {
		byteData, errByteData := json.Marshal(msg)
		if errByteData != nil {
			fmt.Println(errByteData)
		}

		err = ch.Publish(
			"",
			 routingKey,
			false,
			false,
			 amqp.Publishing{
				ContentType: "application/json",
				Body: byteData,
				Priority: 5,
			 })
	}

	failOnErr(err, "Failed to publish a msg")
	log.Printf("Successfully published message %s", data)
}

func PublishIoc(routingKey string, data []model.Indicators) {
	ch, err := ConnectRbmq()
	q, err := ch.QueueDeclare(
		 routingKey,
		true,
		false,
		false,
		false,
		nil,
	)
	fmt.Println(q)
	failOnErr(err, "Failed to create the queue")

	for _, msg := range data {
		byteData, errByteData := json.Marshal(msg)
		if errByteData != nil {
			fmt.Println(errByteData)
		}

		err = ch.Publish(
			"",
			 routingKey,
			false,
			false,
			 amqp.Publishing{
				ContentType: "application/json",
				Body: byteData,
				Priority: 5,
			 })
	}

	failOnErr(err, "Failed to publish a msg")
	log.Printf("Successfully published message %s", data)
}

func PublishSample(routingKey string, data model.Sample) {
	ch, err := ConnectRbmq()
	q, err := ch.QueueDeclare(
		routingKey,
		true,
		false,
		false,
		false,
		nil,
	)
	fmt.Println(q)
	failOnErr(err, "Failed to create the queue")

	byteData, errByteData := json.Marshal(data)
	if errByteData != nil {
		fmt.Println(errByteData)
	}

	err = ch.Publish(
		"",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: byteData,
			Priority: 5,
		})

	failOnErr(err, "Failed to publish a msg")
	log.Printf("Successfully published message %s", data)
}

func PublishCompromised(routingKey string, data []model.Compromised) {
	ch, err := ConnectRbmq()
	q, err := ch.QueueDeclare(
		routingKey,
		true,
		false,
		false,
		false,
		nil,
	)
	fmt.Println(q)
	failOnErr(err, "Failed to create the queue")

	for _, msg := range data {
		byteData, errByteData := json.Marshal(msg)
		if errByteData != nil {
			fmt.Println(errByteData)
		}

		err = ch.Publish(
			"",
			routingKey,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body: byteData,
				Priority: 5,
			})
	}

	failOnErr(err, "Failed to publish a msg")
	log.Printf("Successfully published message %s", data)
}

func failOnErr(err error, msg string) {
	if err != nil {
		fmt.Println("%s: %s", msg, err)
	}
}
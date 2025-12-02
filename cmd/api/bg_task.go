package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"leopard.test/v2/internal/repositories"
)

type Message struct {
	CustomerIds  []int64 `json:"customerIds"`
	CampaignId int64 `json:"campaignId"`
}

func BackgroundTask() {
	rbhost := os.Getenv("MQ_HOST")
	rbport := os.Getenv("MQ_PORT")
	rbusername := os.Getenv("MQ_USERNAME")
	rbpassword := os.Getenv("MQ_PASSWORD")
	rbCampaignQueue := os.Getenv("MQ_CAMPAIGN_QUEUE")
	connecString := fmt.Sprintf("amqp://%s:%s@%s:%s/", rbusername,rbpassword,rbhost, rbport)
	conn, err := amqp.Dial(connecString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	err = ch.Qos(
		1, 
		0,
		false,
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		rbCampaignQueue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Waiting for messages...")

	for m := range msgs {
		log.Printf("Received: %s", m.Body)

		var message Message

		err := json.Unmarshal(m.Body, &message)
		if err != nil {
			log.Println("Failed to parse message:", err)
			m.Nack(false, false) // reject if bad
			continue
		}

		repositories.SendCampaign(message.CustomerIds, message.CampaignId)
		// ack if successfull
		m.Ack(false)
	}
}

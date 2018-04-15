package reciever

import (
	"fluux.io/mqtt"
	"log"
	"github.com/joho/godotenv"
	"os"
	"github.com/funkygao/golib/observer"
)

var address, clientID, topic string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	address = os.Getenv("MQTT_ADDRESS")
	clientID = os.Getenv("CLIENT_ID")
	topic = os.Getenv("TOPIC")
}

func Start() {
	client := mqtt.New(address)
	client.ClientID = clientID
	log.Printf("Server to connect to: %s\n", client.Address)

	messages := make(chan mqtt.Message)
	client.Messages = messages

	postConnect := func(c *mqtt.Client) {
		name := topic
		topic := mqtt.Topic{Name: name, QOS: 1}
		c.Subscribe(topic)
	}

	cm := mqtt.NewClientManager(client, postConnect)
	cm.Start()

	for m := range messages {
		// it was debug info
		//log.Printf("Received message from MQTT server on topic %s: %+v\n", m.Topic, string(m.Payload))
		// here will be publish messages for parser json and websocket publisher
		observer.Publish("dataListener", string(m.Payload))
	}

}
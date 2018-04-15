//
// MQTT basic subscriber
//

package main

import (
	"log"
	"fluux.io/mqtt"
)

var client *mqtt.Client
var messages chan mqtt.Message

func Subscribe() {
	client = mqtt.New("10.11.162.253:1883")
	client.ClientID = "MQTT-Sub"
	log.Printf("Server to connect to: %s\n", client.Address)

	messages = make(chan mqtt.Message)
	client.Messages = messages

	postConnect := func(c *mqtt.Client) {
		name := "orbita_device"
		topic := mqtt.Topic{Name: name, QOS: 1}
		c.Subscribe(topic)
	}

	cm := mqtt.NewClientManager(client, postConnect)
	cm.Start()

	for m := range messages {
		log.Printf("Received message from MQTT server on topic %s: %+v\n", m.Topic, string(m.Payload))
	}
}

func Unsubscribe() {

}

func main() {
	client := mqtt.New("10.11.162.253:1883")
	client.ClientID = "MQTT-Sub"
	log.Printf("Server to connect to: %s\n", client.Address)

	messages := make(chan mqtt.Message)
	client.Messages = messages

	postConnect := func(c *mqtt.Client) {
		name := "orbita_device"
		topic := mqtt.Topic{Name: name, QOS: 1}
		c.Subscribe(topic)
	}

	cm := mqtt.NewClientManager(client, postConnect)
	cm.Start()

	for m := range messages {
		log.Printf("Received message from MQTT server on topic %s: %+v\n", m.Topic, string(m.Payload))
	}
}

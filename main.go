package main

import (
	"log"
	"samsungHackaton/reciever"
	"samsungHackaton/web"
)

func main() {
	log.Println("init application")
	go reciever.Start()
	stopper := make(chan bool, 1)
	go web.StartWsServer()
	go web.StartWebServer()
	<-stopper
}

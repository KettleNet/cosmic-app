package web

import (
	"os"
	"github.com/funkygao/golib/observer"
	"log"
	"github.com/gorilla/websocket"
	"net/http"
	"github.com/joho/godotenv"
)

// module global variables
var websocketAddress string

// initialize stuff
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	websocketAddress = os.Getenv("WS_ADDRESS")
	log.Println("websocket webserverAddress " + webserverAddress)
}

// estabilisher websocket connections and handling
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("connected")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.Println("Upgraded")
	defer c.Close()
	defer func() {
		log.Println("disconnected")
	}()

	closed := make(chan bool, 1)
	go func() {
		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				closed <- true
				return
			}
		}
	}()

	updates := make(chan interface{})
	observer.Subscribe("dataListener", updates)

	for {
		select {
		case update:= <-updates:
			c.WriteMessage(websocket.TextMessage, []byte(update.(string)))
		case <- closed:
			log.Println("Disconnected")
			break
		}

	}
}

func StartWsServer() {
	http.HandleFunc("/echo", handler)
	log.Fatal(http.ListenAndServe(websocketAddress, nil))
}
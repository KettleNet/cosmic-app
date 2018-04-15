package web

import (
	"net/http"
	"log"
	"github.com/joho/godotenv"
	"os"
)

var webserverAddress string

func init()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	webserverAddress = os.Getenv("WEBSERVER_ADDRESS")
	log.Println("web server webserverAddress " + webserverAddress)
}

func StartWebServer()  {
	fs := http.FileServer(http.Dir("./resources"))
	http.Handle("/", fs)
	log.Println("Static webServer launched...")
  	http.ListenAndServe(webserverAddress, nil)
}
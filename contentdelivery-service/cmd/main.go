package main

import (
	"contentDelivery/handlers"
	"contentDelivery/routes"
	"log"
)

const portNumber = ":8000"

func main() {
	r := routes.Routes(handlers.Repo)
	err := r.Run(portNumber)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("contentDelivery-service running on port %s", portNumber)
}

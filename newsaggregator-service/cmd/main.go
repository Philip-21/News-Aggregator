package main

import (
	"log"
	"newsAggregator/handlers"
)

const portNumber = ":8001"

func main() {
	r := Routes(handlers.Repo)
	err := r.Run(portNumber)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("newsAggregator-service running on port %s", portNumber)
}

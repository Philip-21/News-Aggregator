package main

import (
	"log"
	"user/config"
	"user/database"
	"user/handlers"
	"user/routes"
)

const portNumber = ":8080"

func main() {
	conn := database.ConnectToDB()
	if conn == nil {
		log.Panic("Unable to connect to database")
	}
	log.Println("connected to database")
	//app config setup
	app := config.AppConfig{
		DB:     conn,
		Models: database.New(conn),
	}
	r := routes.Routes(handlers.NewRepository(&app))
	err := r.Run(portNumber)
	if err != nil {
		log.Fatal(err)
	}
}

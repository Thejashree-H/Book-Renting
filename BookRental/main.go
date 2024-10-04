package main

import (
	"log"
	"bookrental/config"
	"bookrental/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Starting Books App")

	log.Println("Initializig configuration")
	config := config.InitConfig("books")

	log.Println("Initializig database")
	dbHandler := server.InitDatabase(config)

	log.Println("Initializig HTTP sever")
	httpServer := server.InitHttpServer(config, dbHandler)

	// Start the HTTP server
	httpServer.Start()
}

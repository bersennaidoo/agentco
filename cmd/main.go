package main

import (
	"log"
	"net/http"

	"github.com/bersennaidoo/agentco/application/rest/handlers"
	"github.com/bersennaidoo/agentco/application/rest/server"
	"github.com/bersennaidoo/agentco/infrastructure/repositories/mongo"
	"github.com/bersennaidoo/agentco/physical/config"
	"github.com/bersennaidoo/agentco/physical/dbc"
)

func main() {
	config := config.New(config.GetConfigFileName())
	mclient := dbc.New(config)

	usrepo := mongo.NewUserRepository(mclient)
	hnd := handlers.New(usrepo)
	sgorptions := server.GorillaServerOptions{}
	router := server.HandlerWithOptions(hnd, sgorptions)

	addr := config.GetString("http.http_addr")

	log.Println("Server starting :3000")
	err := http.ListenAndServe(addr, router)
	log.Fatal(err)

}

package main

import (
	"github.com/bersennaidoo/agentco/application/rest/server"
	"github.com/bersennaidoo/agentco/physical/config"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.NewCFGData()

	_ = config.MongoClient(cfg)

	handler := server.Handler{}

	router := gin.Default()

	srvWithOptions := server.GinServerOptions{}

	server.RegisterHandlersWithOptions(router, handler, srvWithOptions)

	router.Run(cfg.Port)
}

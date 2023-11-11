package main

import (
	"github.com/bersennaidoo/agentco/application/rest/server"
	"github.com/gin-gonic/gin"
)

func main() {

	handler := server.Handler{}

	router := gin.Default()

	server.RegisterHandlers(router, handler)

	router.Run(":3000")
}

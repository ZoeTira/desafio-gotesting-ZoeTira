package main

import (
	"github.com/ZoeTira/desafio-gotesting-ZoeTira/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	r.Run(":18085")

}

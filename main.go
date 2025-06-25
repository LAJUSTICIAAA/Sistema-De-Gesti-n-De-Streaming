package main

import (
	"github.com/LAJUSTICIAAA/Sistema-De-Gesti-n-De-Streaming/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	handlers.SetupRoutes(router)
	router.Run(":8080")
}

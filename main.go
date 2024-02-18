package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"telegram-bot-api/config"
	"telegram-bot-api/src/clients"
	"telegram-bot-api/src/handlers"
)

func main() {
	router := gin.Default()
	router.Use(assingCORSconfig())

	bot := clients.Init()
	handlers.Init(bot)

	port := config.Config("PORT")
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(router.Run(":" + port))
}

func assingCORSconfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "*")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

package main

import (
	"log"
	"os"

	"github.com/abhisheksp/heroku-goservice/src/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	router.POST("/testpath/:testparam", handler.TestHandler)

	router.Run(":" + port)
}

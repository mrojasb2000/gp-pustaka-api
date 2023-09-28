package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/hello", helloHandler)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "Mavro",
		"message": "Welcome to Gin Framewor",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content": "Hello World",
		"bio":     "Golang Microservices",
	})
}

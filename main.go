package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/hello", helloHandler)

	router.GET("/books/:id", bookHandler)

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

func bookHandler(c *gin.Context) {
	id, _ := c.Params.Get("id")
	c.JSON(http.StatusOK, gin.H{
		"content": "Books queries",
		"ID":      id,
	})
}

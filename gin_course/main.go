package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var router = gin.Default()
	var address = ":3000"
	router.GET("/", get)
	router.Run(address)
}

func get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
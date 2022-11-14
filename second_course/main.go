package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main()  {
	var router = gin.Default()
	var address = ":3000"
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
			"status": http.StatusOK,
		})
	})
	router.Run(address)
}
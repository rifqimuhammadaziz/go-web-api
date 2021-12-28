package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    "Rifqi Muhammad Aziz",
			"address": "Tegal, Central Java",
			"bio":     "Software Engineer",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"content":     "Hello World",
			"description": "Ini adalah hello world",
		})
	})

	router.Run()
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)

	// change listening port
	router.Run(":8000")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "Rifqi Muhammad Aziz",
		"address": "Tegal, Central Java",
		"bio":     "Software Engineer",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":       "Hello World",
		"description": "Ini adalah hello world",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")       // get id from url parameter (localhost:8000/books/2)
	title := c.Param("title") // get title from url parameter (localhost:8000/books/2/ini-adalah-judul)
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(c *gin.Context) {
	// localhost:8000/query?price=40&title=ini adalah judul buku
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"price": price,
		"title": title,
	})
}

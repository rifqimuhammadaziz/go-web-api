package main

import (
	"fmt"
	"log"
	"myweb-api/book"
	"myweb-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// https://github.com/go-sql-driver/mysql
	dsn := "root:root@tcp(localhost:3306)/myweb_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	// migrate struct to database
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	books, err := bookRepository.FindAll()
	for _, book := range books {
		fmt.Println(book)
	}

	router := gin.Default()

	// API Versioning
	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)

	v1.POST("/books", handler.PostBooksHandler)

	// change listening port
	router.Run()
}

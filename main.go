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

	/**
	CREATE DATA
	create data every running server.
	*/
	// book := book.Book{}
	// book.Title = "Buku kedua"
	// book.Price = 120000
	// book.Rating = 5
	// book.Description = "Ini adalah deskripsi buku kedua"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("=========================")
	// 	fmt.Println("Failed create book record")
	// 	fmt.Println("=========================")
	// }

	/**
	READ DATA
	*/
	var books []book.Book

	// run & print query or error
	err = db.Debug().Where("id = ?", 1).Find(&books).Error
	if err != nil {
		fmt.Println("=========================")
		fmt.Println("Failed find book record")
		fmt.Println("=========================")
	}

	// get data by query
	for _, book := range books {
		fmt.Println("Title: ", book.Title)
		fmt.Printf("Book Object: %v", book)
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

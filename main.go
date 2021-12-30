package main

import (
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

	// bookRepository := book.NewRepository(db)
	bookFileRepository := book.NewFileRepository()

	bookService := book.NewService(bookFileRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	// API Versioning
	v1 := router.Group("/v1")
	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	v1.GET("/books/:id/:title", bookHandler.BooksHandler)
	v1.GET("/query", bookHandler.QueryHandler)

	v1.POST("/books", bookHandler.PostBooksHandler)

	// change listening port
	router.Run()

	//main
	//handler
	//service
	//repository
	//db
	//mysql
}

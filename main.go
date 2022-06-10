package main

import (
	"log"
	"rest-api-golang/book"
	"rest-api-golang/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	router := gin.Default()

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)

	book := book.Book{
		Title:       "Buku Masak",
		Description: "Masak itu asik loh",
		Price:       90000,
		Rating:      4,
		Discount:    0,
	}

	bookRepository.Create(book)

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/about", handler.AboutHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("books", handler.PostBooksHandler)

	router.Run()
}

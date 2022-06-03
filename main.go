package main

import (
	"fmt"
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
	// CRUD

	// CREATE data
	// book := book.Book{}
	// book.Title = "Ngoding itu asik"
	// book.Price = 30000
	// book.Rating = 4
	// book.Discount = 15
	// book.Description = "Buku yang mengulas tentang programming dari basic sampai menengah"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Creating Book Record")
	// 	fmt.Println("==========================")
	// }

	// Read data
	var books []book.Book
	err = db.Debug().Where("id = ?", 1).First(&books).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Error finding Book Record")
		fmt.Println("==========================")
	}
	for _, b := range books {
		fmt.Println("Title : ", b.Title)
		fmt.Printf("Book Object %v", b)
	}

	// Update Data
	var book book.Book
	err = db.Debug().Where("id = ?", 1).First(&book).Error

	book.Title = "Man Ringer (Revised edition)"
	err = db.Save(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Error updating Book Record")
		fmt.Println("==========================")
	}

	// Delete Data
	err = db.Debug().Where("id = ?", 1).First(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Error finding Book Record")
		fmt.Println("==========================")
	}

	err = db.Delete(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Error deleting Book Record")
		fmt.Println("==========================")
	}

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/about", handler.AboutHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("books", handler.PostBooksHandler)

	router.Run()
}

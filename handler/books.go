package handler

import (
	"fmt"
	"net/http"

	"rest-api-golang/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Agrieva Xananda",
		"bio":  "A junior Software Developers and Game Developer",
	})
}

func AboutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content": "Anak umur 16 tahun",
		"motto":   "Keep Learning",
	})
}

// mengambil variable id memakai .Param()
func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

// jika memakai query string maka akan memakai .Query()
func QueryHandler(c *gin.Context) {
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{"title": title})
}

func PostBooksHandler(c *gin.Context) {
	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {

		errorMassages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := fmt.Sprintf("error in filed %s, condition: %s,", e.Field(), e.ActualTag())
			errorMassages = append(errorMassages, errorMassage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMassages,
		})
		return

	} else {
		c.JSON(http.StatusOK, gin.H{
			"title": bookInput.Title,
			"price": bookInput.Price,
		})
	}
}

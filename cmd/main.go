package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	books_data := map[string]map[string]string{
		"B001": map[string]string{
			"name": "Book 1",
			"owner": "U001",
		},
		"B002": map[string]string{
			"name": "Book 2",
			"owner": "U002",
		},
		"B003": map[string]string{
			"name": "Book 3",
			"owner": "U003",
		},
	}
	
	router := gin.Default()

	router.GET("/api/v1/books", func(c *gin.Context) {
		c.JSON(200, books_data)
	})

	router.Run(":8000")
}


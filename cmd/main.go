package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func main() {
	books_data := map[string]map[string]string{
		"B001": map[string]string{
			"name": "Book 1",
			"owner": "U001",
			"status": "AVAILABLE",
		},
		"B002": map[string]string{
			"name": "Book 2",
			"owner": "U002",
			"status": "AVAILABLE",
		},
		"B003": map[string]string{
			"name": "Book 3",
			"owner": "U003",
			"status": "AVAILABLE AFTER 2020-01-03 00:00:00",
		},
	}
	borrow_data := map[string]map[string]string{
		"R001": map[string]string{
			"book_id": "B003",
			"borrower": "U004",
			"borrow_start_time": "2020-01-01 00:00:00",
			"borrow_end_time": "2020-01-03 00:00:00",
		},
	}

	router := gin.Default()

	router.GET("/api/v1/books", func(c *gin.Context) {
		c.JSON(200, books_data)
	})

	router.GET("/api/v1/books/:id", func(c *gin.Context) {
		book_id := c.Param("id")
		c.JSON(200, books_data[book_id])
	})

	router.PUT("/api/v1/books", func(c *gin.Context) {
		var json map[string]string
		c.BindJSON(&json)

		book_id := strconv.Itoa(len(books_data) + 1)
		for len(book_id) < 3 {
			book_id = "0" + book_id
		}

		book_id = "B" + book_id

		books_data[book_id] = json
		books_data[book_id]["status"] = "AVAILABLE"
		
		c.JSON(200, gin.H{
			"book_id": book_id,
		})
	})

	router.GET("/api/v1/books/borrowed", func(c *gin.Context) {
		c.JSON(200, borrow_data)
	})

	router.GET("/api/v1/books/borrowed/:id", func(c *gin.Context) {
		borrow_id := c.Param("id")
		c.JSON(200, borrow_data[borrow_id])
	})

	router.PUT("/api/v1/books/:id/borrow", func(c *gin.Context) {
		book_id := c.Param("id")
		var json map[string]string
		c.BindJSON(&json)

		borrow_id := strconv.Itoa(len(borrow_data) + 1)
		for len(borrow_id) < 3 {
			borrow_id = "0" + borrow_id
		}

		borrow_id = "R" + borrow_id

		borrow_data[borrow_id] = map[string]string{
			"book_id": book_id,
			"borrower": json["borrower"],
			"borrow_start_time": json["borrow_start_time"],
			"borrow_end_time": json["borrow_end_time"],
		}

		books_data[book_id]["status"] = "AVAILABLE AFTER " + json["borrow_end_time"]

		c.JSON(200, gin.H{
			"borrow_id": borrow_id,
		})
	})

	router.POST("/api/v1/books/borrowed/:id/return", func(c *gin.Context) {
		borrow_id := c.Param("id")
		book_id := borrow_data[borrow_id]["book_id"]

		delete(borrow_data, borrow_id)

		books_data[book_id]["status"] = "AVAILABLE"

		c.JSON(200, gin.H{
			"book_id": book_id,
		})
	})

	router.Run(":8000")
}


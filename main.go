package main

import (
	"net/http"

	"gomark/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// connect to postgres database
	database.Connect()

	// initialize gin router
	r := gin.Default()

	// test route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "GoMark API - Smart Bookmark Manager"})
	})

	// run the server
	r.Run(":8080")
}

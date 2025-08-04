package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize gin router
	r := gin.Default()

	// test route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "GoMark API - Smart Bookmark Manager"})
	})

	// run the server
	r.Run(":8080")
}

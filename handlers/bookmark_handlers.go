package handlers

import (
	"gomark/database"
	"gomark/models"
	"gomark/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create a new bookmark
func CreateBookmark(c *gin.Context) {
	log.Printf("CreateBookmark handler called")
	var bookmark models.Bookmark

	// bind JSON from request body
	if err := c.ShouldBindJSON(&bookmark); err != nil {
		log.Printf("JSON binding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// scrape the URL for title and description if not provided
	log.Printf("Bookmark title: '%s', description: '%s'", bookmark.Title, bookmark.Description)
	if bookmark.Title == "" || bookmark.Description == "" {
		log.Printf("Attempting to scrape URL: %s", bookmark.URL)
		scraped := services.ScrapeURL(bookmark.URL)

		if scraped.Error == nil {
			if bookmark.Title == "" {
				bookmark.Title = scraped.Title
			}
			if bookmark.Description == "" {
				bookmark.Description = scraped.Description
			}
		} else {
			log.Printf("Scraping failed: %v", scraped.Error)
		}
	} else {
		log.Printf("Skipping scraping - title and description already provided")
	}

	// save to database
	if err := database.DB.Create(&bookmark).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save bookmark"})
		return
	}

	c.JSON(http.StatusCreated, bookmark)
}

// get all bookmarks
func GetBookmarks(c *gin.Context) {
	var bookmarks []models.Bookmark

	if err := database.DB.Find(&bookmarks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookmarks"})
		return
	}

	c.JSON(http.StatusOK, bookmarks)
}

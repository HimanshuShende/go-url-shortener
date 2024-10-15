package main

import (
	"fmt"

	"github.com/HimanshuShende/go-url-shortner/handler"
	"github.com/HimanshuShende/go-url-shortner/store"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	router.POST("/create-short-url", func(ctx *gin.Context) {
		handler.CreateShortUrl(ctx)
	})

	router.GET("/:shortUrl", func(ctx *gin.Context) {
		handler.HandleShortIrlRedirect(ctx)
	})

	// Store initialization
	store.InitializeStore()

	err := router.Run(":9808")

	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}

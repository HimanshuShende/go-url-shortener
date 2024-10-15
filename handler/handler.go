package handler

import (
	"net/http"

	"github.com/HimanshuShende/go-url-shortner/shortener"
	"github.com/HimanshuShende/go-url-shortner/store"
	"github.com/gin-gonic/gin"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(ctx *gin.Context) {
	var creationReq UrlCreationRequest
	if err := ctx.ShouldBindJSON(&creationReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortUrl := shortener.GenerateShortUrl(creationReq.LongUrl, creationReq.UserId)
	store.SaveUrlMapping(shortUrl, creationReq.LongUrl, creationReq.UserId)

	host := "http://localhost:9808/"
	ctx.JSON(200, gin.H{
		"message":   "Short URL created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortIrlRedirect(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	ctx.Redirect(302, initialUrl)
}

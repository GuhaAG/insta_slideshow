package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guhaag/insta_slideshow/backend/instagram"
)

func V1(router *gin.RouterGroup) {
	router.GET("/images", func(context *gin.Context) {
		images := instagram.ScrapeMedia()
		context.JSON(http.StatusOK, images)
	})
}

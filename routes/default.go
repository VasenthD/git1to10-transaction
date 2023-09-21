package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Default(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Server is Healthy"})
	})
}

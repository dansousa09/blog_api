package handler

import "github.com/gin-gonic/gin"

func CreatePostHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "CreatePostHandler",
	})
}

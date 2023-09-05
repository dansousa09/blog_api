package handler

import "github.com/gin-gonic/gin"

func ShowPostHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ShowPostHandler",
	})
}

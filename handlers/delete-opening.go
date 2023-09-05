package handler

import "github.com/gin-gonic/gin"

func DeletePostHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "DeletePostHandler",
	})
}

package handler

import "github.com/gin-gonic/gin"

func UpdatePostHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "UpdatePostHandler",
	})
}

package handler

import "github.com/gin-gonic/gin"

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "DeleteOpeningHandler",
	})
}

package handler

import "github.com/gin-gonic/gin"

func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ListOpeningsHandler",
	})
}

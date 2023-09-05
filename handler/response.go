package handler

import "github.com/gin-gonic/gin"

func sendErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"error": message,
	})
}

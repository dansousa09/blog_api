package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"error": message,
	})
}

func sendSuccessResponse(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, gin.H{
		"message": fmt.Sprintf("operation %s was successfully completed", operation),
		"data":    data,
	})
}

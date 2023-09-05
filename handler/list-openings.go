package handler

import "github.com/gin-gonic/gin"

func ListPostsHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ListPostsHandler",
	})
}

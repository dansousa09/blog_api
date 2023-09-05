package handler

import (
	"net/http"

	"github.com/dansousa09/blog_api/schemas"
	"github.com/gin-gonic/gin"
)

func ListPostsHandler(ctx *gin.Context) {
	posts := []schemas.Post{}

	if err := db.Find(&posts).Error; err != nil {
		logger.Errorf("error on list posts: %s", err.Error())
		sendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(ctx, "list posts", posts)
}

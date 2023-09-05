package handler

import (
	"net/http"

	"github.com/dansousa09/blog_api/schemas"
	"github.com/gin-gonic/gin"
)

func CreatePostHandler(ctx *gin.Context) {
	request := CreatePostRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %s", err.Error())
		sendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	post := schemas.Post{
		Title:   request.Title,
		Content: request.Content,
		Author:  request.Author,
	}

	if err := db.Create(&post).Error; err != nil {
		logger.Errorf("error on create post: %s", err.Error())
		sendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(ctx, "create post", post)
}

package handler

import (
	"net/http"

	"github.com/dansousa09/blog_api/schemas"
	"github.com/gin-gonic/gin"
)

const errUpdatePostErrorMsg = "error on update post"

func UpdatePostHandler(ctx *gin.Context) {
	request := UpdatePostRequest{}

	ctx.BindJSON(request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error on validate post: %s", err)
		sendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		err := errFieldIsRequired("id", "query-param")
		logger.Errorf("error on get post: %s", err)
		sendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	post := schemas.Post{}

	if err := db.First(&post, id).Error; err != nil {
		logger.Errorf("error on get post: %s", err)
		sendErrorResponse(ctx, http.StatusNotFound, errPostNotFoundMsg)
		return
	}

	if request.Title != "" {
		post.Title = request.Title
	}

	if request.Content != "" {
		post.Content = request.Content
	}

	if request.Author != "" {
		post.Author = request.Author
	}

	if err := db.Save(&post).Error; err != nil {
		logger.Errorf("%s: %s", errUpdatePostErrorMsg, err)
		sendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(ctx, "update post", post)
}

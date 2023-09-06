package handler

import (
	"net/http"

	"github.com/dansousa09/blog_api/schemas"
	"github.com/gin-gonic/gin"
)

const (
	errFieldIsRequiredMsg = "field is required"
	errPostNotFoundMsg    = "post not found"
	errDeletePostErrorMsg = "error on delete post"
)

func DeletePostHandler(ctx *gin.Context) {
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

	if err := db.Delete(&post).Error; err != nil {
		logger.Errorf("%s: %s", errDeletePostErrorMsg, err)
		sendErrorResponse(ctx, http.StatusInternalServerError, errDeletePostErrorMsg)
		return
	}

	sendSuccessResponse(ctx, "delete post", post)
}

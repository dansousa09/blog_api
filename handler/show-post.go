package handler

import (
	"net/http"

	"github.com/dansousa09/blog_api/schemas"
	"github.com/gin-gonic/gin"
)

func ShowPostHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		logger.Errorf("error on show post: %s", errFieldIsRequired("id", "query-param"))
		sendErrorResponse(ctx, http.StatusBadRequest, errFieldIsRequired("id", "query-param").Error())
		return
	}

	post := schemas.Post{}

	if err := db.First(&post, id).Error; err != nil {
		logger.Error("error on show post: post not found")
		sendErrorResponse(ctx, http.StatusNotFound, "post not found")
		return
	}

	sendSuccessResponse(ctx, "show post", post)
}

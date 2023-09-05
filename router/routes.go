package router

import (
	handler "github.com/dansousa09/blog_api/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	basePath := "/api/v1"

	v1 := r.Group(basePath)
	{
		v1.GET("/post", handler.ShowPostHandler)
		v1.GET("/posts", handler.ListPostsHandler)
		v1.POST("/post", handler.CreatePostHandler)
		v1.PUT("/post", handler.UpdatePostHandler)
		v1.DELETE("/post", handler.DeletePostHandler)
	}
}

package router

import (
	handler "github.com/dansousa09/blog_api/handlers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	basePath := "/api/v1"

	v1 := r.Group(basePath)
	{
		v1.GET("/opening", handler.CreateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}
}

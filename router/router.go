package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()

	InitializeRoutes(r)

	r.Run(":8080")
}

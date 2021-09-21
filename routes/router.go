package routes

import (
	middleware "BiliBili.com/midderware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.Group("/api/v1")
	return r
}

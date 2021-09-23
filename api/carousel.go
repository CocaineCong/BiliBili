package api

import (
	"BiliBili.com/service"
	"github.com/gin-gonic/gin"
)

func Carousel(c *gin.Context) {
	var carouselService service.Carousel
	_ = c.ShouldBind(&carouselService)
	res := carouselService.Carousel()
	c.JSON(200,res)
}
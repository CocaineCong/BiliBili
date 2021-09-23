package api

import (
	"BiliBili.com/service"
	"github.com/gin-gonic/gin"
)

func VideoShow(c *gin.Context) {
	var videoShowService service.VideoShow
	_ = c.ShouldBind(&videoShowService)
	res := videoShowService.Show(c.Param("id"))
	c.JSON(200,res)
}

func VideoRecommend(c *gin.Context) {
	var videoRecommend service.VideoRecommend
	_ = c.ShouldBind(&videoRecommend)
	res := videoRecommend.Recommend()
	c.JSON(200,res)
}

func VideoList(c *gin.Context) {
	var videoListService service.VideoShow
	_ = c.ShouldBind(&videoListService)
	res := videoListService.List(c.Param("id"))
	c.JSON(200,res)
}

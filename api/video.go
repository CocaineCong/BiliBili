package api

import (
	"BiliBili.com/pkg/utils"
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

func VideoFavorite(c *gin.Context) {
	var videoFavoriteService service.VideoShow
	_ = c.ShouldBind(&videoFavoriteService)
	res := videoFavoriteService.Favor(c.Param("id"))
	c.JSON(200,res)
}

func VideoUpdate(c *gin.Context) {
	var videoUpdateService service.VideoInfo
	_ = c.ShouldBind(&videoUpdateService)
	_,chaim,_ := utils.ParseUserToken(c.GetHeader("Authorization"))
	res:=videoUpdateService.Update(c.Param("id"),chaim.UserId)
	c.JSON(200,res)
}

func VideoDelete(c *gin.Context) {
	var videoDeleteService service.VideoDelete
	_ = c.ShouldBind(&videoDeleteService)
	res := videoDeleteService.Delete(c.Param("id"))
	c.JSON(200,res)
}

func VideoUpload(c *gin.Context) {
	var videoUploadService service.VideoShow
	file , fileHeader  ,_ := c.Request.FormFile("video")
	cover , coverHeader  ,_ := c.Request.FormFile("cover")
	_ = c.ShouldBind(&videoUploadService)
	_,chaim,_ := utils.ParseUserToken(c.GetHeader("Authorization"))
	res:=videoUploadService.Upload(chaim.UserId,file,cover,coverHeader.Size,fileHeader.Size)
	c.JSON(200,res)
}

func VideoInteractiveData(c *gin.Context) {
	var videoInteractiveService service.VideoInteractiveData
	_ = c.ShouldBind(&videoInteractiveService)
	_,chaim,_ := utils.ParseUserToken(c.Param("Authorization"))
	res := videoInteractiveService.Show(c.Param("id"),chaim.UserId)
	c.JSON(200,res)
}
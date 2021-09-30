package api

import (
	"BiliBili.com/pkg/utils"
	"BiliBili.com/service"
	"github.com/gin-gonic/gin"
)

func FavorCreate(c *gin.Context) {
	var favorCreate service.FavorCreateService
	_ = c.ShouldBind(&favorCreate)
	_,chain,_ := utils.ParseUserToken(c.GetHeader("Authorization"))
	res:=favorCreate.Create(c.Param("id"),chain.UserId)
	c.JSON(200,res)
}

func FavorDelete(c *gin.Context) {
	var favorDelete service.FavorDeleteService
	_ = c.ShouldBind(&favorDelete)
	_,chain,_ := utils.ParseUserToken(c.GetHeader("Authorization"))
	res:=favorDelete.Delete(c.Param("id"),chain.UserId)
	c.JSON(200,res)
}

func Like(c *gin.Context) {
	var likeCreate service.LikeCreateService
	_ = c.ShouldBind(&likeCreate)
	_,chain,_ := utils.ParseUserToken(c.GetHeader("Authorization"))
	res:=likeCreate.Create(c.Param("id"),chain.UserId)
	c.JSON(200,res)
}

func Dislike(c *gin.Context) {
	var likeDelete service.LikeDeleteService
	_ = c.ShouldBind(&likeDelete)
	_,chain,_ := utils.ParseUserToken(c.GetHeader("Authorization"))
	res:=likeDelete.Delete(c.Param("id"),chain.UserId)
	c.JSON(200,res)
}
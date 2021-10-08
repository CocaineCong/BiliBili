package api

import (
	"BiliBili.com/service"
	"github.com/gin-gonic/gin"
)

func ListFollowing(c *gin.Context) {  //关注列表
	var listFollowingService service.ListFollowingService
	_ = c.ShouldBind(&listFollowingService)
	res := listFollowingService.List(c.Param("id"))
	c.JSON(200,res)
}

func ListFollower(c *gin.Context) {  //粉丝列表
	var listFollowerService service.ListFollowerService
	_ = c.ShouldBind(&listFollowerService)
	res := listFollowerService.List(c.Param("id"))
	c.JSON(200,res)
}
package api

import (
	"BiliBili.com/pkg/utils"
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

func CreateFollowing(c *gin.Context) {
	var createFollowingService service.CreateFollowingService
	_ = c.ShouldBind(&createFollowingService)
	_,chaim,_ := utils.ParseUserToken(c.GetHeader("Authorization"))
	res := createFollowingService.Create(c.Param("id"),chaim.UserId) // 关注谁就是谁的id
	c.JSON(200,res)
}

func DeleteFollowing(c *gin.Context) {
	var deleteFollowingService service.DeleteFollowingService
	_ = c.ShouldBind(&deleteFollowingService)
	_,chaim,_ := utils.ParseUserToken(c.GetHeader("Authorization"))
	res := deleteFollowingService.Delete(c.Param("id"),chaim.UserId) // 关注谁就是谁的id
	c.JSON(200,res)
}

package api

import (
	"BiliBili.com/pkg/utils"
	"BiliBili.com/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegisterService service.UserRegister
	_ = c.ShouldBind(&userRegisterService)
	res := userRegisterService.Register()
	c.JSON(200,res)
}

func UserLogin(c *gin.Context) {
	var userLoginService service.UserLogin
	_ = c.ShouldBind(&userLoginService)
	res := userLoginService.Login()
	c.JSON(200,res)
}

func UserUpdate(c *gin.Context) {
	var userUpdateService service.UserUpdate
	_ = c.ShouldBind(&userUpdateService)
	_,chaim,_ := utils.ParseUserToken(c.GetHeader("Authorization"))
	res := userUpdateService.Update(chaim.UserId)
	c.JSON(200,res)
}

func UserInfo(c *gin.Context) {
	var userInfoService service.UserInfo
	_ = c.ShouldBind(&userInfoService)
	_,chaim,_ := utils.ParseUserToken(c.GetHeader("Authorization"))
	res := userInfoService.Show(chaim.UserId)
	c.JSON(200,res)
}

func UserSearch(c *gin.Context) {
	var userSearchService service.UserSearch
	_ = c.ShouldBind(&userSearchService)
	res := userSearchService.Search()
	c.JSON(200,res)
}


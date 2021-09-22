package api

import (
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
	res := userLoginService.UserLogin()
	c.JSON(200,res)
}
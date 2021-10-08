package api

import (
	"BiliBili.com/service"
	"github.com/gin-gonic/gin"
)

func ListDanmu(c *gin.Context) {
	var listDanmuService service.ListDanmuService
	_ = c.ShouldBind(&listDanmuService)
	res := listDanmuService.List(c.Param("vid"))
	c.JSON(200,res)
}

func CreateDanmu(c *gin.Context) {
	var createDanmuService service.CreateDamuService
	_ = c.ShouldBind(&createDanmuService)
	res := createDanmuService.Create(c.Param("uid"))
	c.JSON(200,res)
}

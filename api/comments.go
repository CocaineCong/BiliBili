package api

import (
	"BiliBili.com/service"
	"github.com/gin-gonic/gin"
)

func CommentsGet(c *gin.Context) {
	var commentService service.ShowCommentService
	_ = c.ShouldBind(&commentService)
	res:=commentService.Show(c.Param("id"))
	c.JSON(200,res)
}



package api

import (
	"BiliBili.com/pkg/utils"
	"BiliBili.com/service"
	"github.com/gin-gonic/gin"
)

func CommentsGet(c *gin.Context) {
	var commentShowService service.ShowCommentService
	_ = c.ShouldBind(&commentShowService)
	res:=commentShowService.Show(c.Param("id"))
	c.JSON(200,res)
}

func CommentsDelete(c *gin.Context) {
	var deleteCommentService service.DeleteCommentService
	_ = c.ShouldBind(&deleteCommentService)
	_,chaim,_ := utils.ParseUserToken(c.GetHeader("Authorization"))
	res := deleteCommentService.Delete(c.Param("id"),chaim.UserId)
	c.JSON(200,res)
}

func ReplyDelete(c *gin.Context) {
	var deleteReplyService service.DeleteReplyService
	_ = c.ShouldBind(&deleteReplyService)
	_,chaim,_:=utils.ParseUserToken(c.GetHeader("Authorization"))
	res := deleteReplyService.Delete(c.Param("id"),chaim.UserId)
	c.JSON(200,res)
}

func CreateComment(c *gin.Context) {
	var createCommentService service.CreateCommentService
	_,chaim,_:=utils.ParseUserToken(c.GetHeader("Authorization"))
	res:=createCommentService.Create(c.Param("vid"),chaim.UserId)
	c.JSON(200,res)
}

func CreateReply(c *gin.Context) {
	var createReplyService service.CreateReplyService
	_,chaim,_:=utils.ParseUserToken(c.GetHeader("Authorization"))
	res:=createReplyService.Create(c.Param("cid"),chaim.UserId)
	c.JSON(200,res)
}
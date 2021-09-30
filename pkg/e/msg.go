package e

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",

	ErrorAuthCheckTokenFail:        "Token验证失败",
	ErrorAuthCheckTokenTimeout:     "Token股哦其",
	ErrorAuthInsufficientAuthority: "无权限",

	ErrorUploadFile : "上传文件失败",

	ErrorFavorExist :"已经收藏了",
	ErrorLikeExist :"已经点赞了",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
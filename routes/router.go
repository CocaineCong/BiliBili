package routes

import (
	"BiliBili.com/api"
	middleware "BiliBili.com/midderware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())
	v1 := r.Group("api/v1")
	{
		// 用户操作
		v1.POST("user/register",api.UserRegister)
		v1.POST("user/login",api.UserLogin)

		// 轮播图
		v1.GET("carousels",api.Carousel)

		//视频操作
		v1.GET("video/:id",api.VideoShow)
		v1.GET("video-recommend", api.VideoRecommend)

		authed := v1.Group("/")            //需要登陆保护
		authed.Use(middleware.JWT())
		{
			authed.PUT("user/update",api.UserUpdate)
			authed.GET("user/show",api.UserInfo)
			authed.POST("user/search",api.UserSearch)

			authed.GET("video/:id",api.VideoList)
			authed.GET("video-favor/:id",api.VideoFavorite)
		}
	}
	return r
}

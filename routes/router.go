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

		v1.GET("comment/:id",api.CommentsGet)
		authed := v1.Group("/")            //需要登陆保护
		authed.Use(middleware.JWT())
		{
			//用户操作
			authed.PUT("user/update",api.UserUpdate)
			authed.GET("user/show",api.UserInfo)
			authed.POST("user/search",api.UserSearch)

			//视频操作
			authed.GET("video/:id",api.VideoList)
			authed.GET("video-favor/:id",api.VideoFavorite)
			authed.PUT("video/:id",api.VideoUpdate)
			authed.DELETE("video/:id",api.VideoDelete)
			authed.POST("video",api.VideoUpload)

			//互动操作
			authed.GET("/video-interactive/:id", api.VideoInteractiveData) //获取点赞收藏关注的交互数据
			authed.POST("/favor/:id", api.FavorCreate)
			authed.DELETE("/favor/:id", api.FavorDelete)
			authed.POST("/like/:id", api.Like)
			authed.DELETE("/like/:id", api.Dislike)

			//评论回复

		}
	}
	return r
}

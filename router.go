package main

import (
	"mydousheng/controller"
	"mydousheng/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter()  {
	r := gin.Default()
	
	douyin := r.Group("/douyin")
	// basic
	{
		user := douyin.Group("/user")
		{
			user.POST("/register/",controller.Register)
			user.POST("/login/",controller.Login)
			user.GET("/",middleware.AuthMiddleWare(),controller.UserInfo)
		}
		publish := douyin.Group("/publish")
		{
			publish.POST("/action/",middleware.AuthMiddleWare(),controller.PublishAction)
			publish.GET("/list/",middleware.AuthMiddleWare(),controller.PublishList)
		}
		feed := douyin.Group("/feed")
		{
			feed.GET("/",controller.Feed)
		}
	}
	

	r.Run(":8081")
	

	
}
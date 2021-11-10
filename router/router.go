package router

import (
	"my-gram/controller"
	"my-gram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.UserLogin)
		userRouter.PUT("/", middlewares.Authentication(), controller.UserUpdate)
		userRouter.DELETE("/", middlewares.Authentication(), controller.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controller.PostPhoto)
		photoRouter.GET("/", controller.ReadAllPhoto)
		photoRouter.PUT(":photoId", controller.IndexHandler)
		photoRouter.DELETE(":photoId", controller.IndexHandler)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.POST("/", controller.IndexHandler)
		commentRouter.GET("/", controller.IndexHandler)
		commentRouter.PUT(":commentId", controller.IndexHandler)
		commentRouter.DELETE(":commentId", controller.IndexHandler)
	}

	socialmediaRouter := r.Group("/socialmedias")
	{
		socialmediaRouter.POST("/", controller.IndexHandler)
		socialmediaRouter.GET("/", controller.IndexHandler)
		socialmediaRouter.PUT(":socialMediaId", controller.IndexHandler)
		socialmediaRouter.DELETE(":socialMediaId", controller.IndexHandler)
	}
	return r
}

package router

import (
	"my-gram/controller"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.IndexHandler)
		userRouter.PUT("/", controller.IndexHandler)
		userRouter.DELETE("/", controller.IndexHandler)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.POST("/", controller.IndexHandler)
		photoRouter.GET("/", controller.IndexHandler)
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
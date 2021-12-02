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
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controller.StoreComment)
		commentRouter.GET("/", controller.GetComments)
		commentRouter.PUT(":commentId", controller.UpdateComment)
		commentRouter.DELETE(":commentId", controller.DeleteComment)
	}

	socialmediaRouter := r.Group("/socialmedias")
	{
		socialmediaRouter.Use(middlewares.Authentication())
		socialmediaRouter.POST("/", controller.StoreSocialMedia)
		socialmediaRouter.GET("/", controller.GetSocialMedias)
		socialmediaRouter.PUT(":socialMediaId", controller.UpdateSocialMedia)
		socialmediaRouter.DELETE(":socialMediaId", controller.DeleteSocialMedia)
	}
	return r
}

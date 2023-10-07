package router

import (
	"basicGin/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter() {
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublicUser := rgPublic.Group("user")
		{
			rgPublicUser.POST("/login", userApi.Login)
		}

		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.POST("", userApi.AddUser)
			rgAuthUser.GET("/:id", userApi.GetUserById)
			rgAuthUser.GET("", userApi.GetUserList)
			rgAuthUser.PUT("/:id", userApi.UpdateUser)
			rgAuthUser.DELETE("/:id", userApi.DeleteUserById)
		}

	})
}

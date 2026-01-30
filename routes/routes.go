package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/knr1997/quiz-tracker-apiserver/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/users", controllers.CreateUser)
		api.GET("/users", controllers.GetUsers)
		api.POST("/getUserList", baseApi.GetUserList)
	}
}

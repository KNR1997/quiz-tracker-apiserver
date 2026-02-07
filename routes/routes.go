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
		api.POST("/customers", exaCustomerApi.CreateExaCustomer)

		api.POST("/courses", courseApi.CreateCourse)
		api.PUT("/courses/:id", courseApi.UpdateCourse)
		api.GET("/courses", courseApi.GetCourseList)
		api.DELETE("/courses/:id", courseApi.DeleteCourse)

		api.GET("/quizzes", quizApi.GetQuizList)
		api.POST("/quizzes", quizApi.CreateQuiz)
		api.PUT("/quizzes/:id", quizApi.UpdateQuiz)
		api.DELETE("/quizzes/:id", quizApi.DeleteQuiz)

	}
}

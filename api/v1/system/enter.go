package system

import "github.com/knr1997/quiz-tracker-apiserver/service"

type ApiGroup struct {
	BaseApi
	CourseApi
	QuizApi
}

var (
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	courseService = service.ServiceGroupApp.SystemServiceGroup.CourseService
	quizService   = service.ServiceGroupApp.SystemServiceGroup.QuizService
)

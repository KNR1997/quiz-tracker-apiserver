package routes

import api "github.com/knr1997/quiz-tracker-apiserver/api/v1"

type RouterGroup struct {
}

var (
	exaCustomerApi = api.ApiGroupApp.ExampleApiGroup.CustomerApi
	baseApi        = api.ApiGroupApp.SystemApiGroup.BaseApi
	courseApi      = api.ApiGroupApp.SystemApiGroup.CourseApi
	quizApi        = api.ApiGroupApp.SystemApiGroup.QuizApi
)

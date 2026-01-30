package system

import "github.com/knr1997/quiz-tracker-apiserver/service"

type ApiGroup struct {
	BaseApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)

package example

import "github.com/knr1997/quiz-tracker-apiserver/service"

type ApiGroup struct {
	CustomerApi
}

var (
	customerService = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
)

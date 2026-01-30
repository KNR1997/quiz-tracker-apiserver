package service

import (
	"github.com/knr1997/quiz-tracker-apiserver/service/example"
	"github.com/knr1997/quiz-tracker-apiserver/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

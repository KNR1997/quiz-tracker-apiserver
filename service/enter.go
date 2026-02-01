package service

import (
	"github.com/knr1997/quiz-tracker-apiserver/service/example"
	"github.com/knr1997/quiz-tracker-apiserver/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

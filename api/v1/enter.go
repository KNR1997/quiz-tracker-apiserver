package v1

import (
	"github.com/knr1997/quiz-tracker-apiserver/api/v1/example"
	"github.com/knr1997/quiz-tracker-apiserver/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

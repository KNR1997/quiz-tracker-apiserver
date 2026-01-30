package v1

import (
	"github.com/knr1997/quiz-tracker-apiserver/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

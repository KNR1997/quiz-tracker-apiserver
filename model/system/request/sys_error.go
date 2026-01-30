package request

import (
	"time"

	"github.com/knr1997/quiz-tracker-apiserver/model/common/request"
)

type SysErrorSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Form           *string     `json:"form" form:"form"`
	Info           *string     `json:"info" form:"info"`
	request.PageInfo
}

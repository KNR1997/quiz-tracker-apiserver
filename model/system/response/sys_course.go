package response

import "github.com/knr1997/quiz-tracker-apiserver/model/system"

type CourseResponse struct {
	Course system.SysCourse `json:"course"`
}

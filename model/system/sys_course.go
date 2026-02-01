package system

import "github.com/knr1997/quiz-tracker-apiserver/global"

type SysCourse struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"comment:Course Name"`
	Slug string `json:"slug" form:"slug" gorm:"comment:Course Slug"`
}

package system

import (
	"time"

	"github.com/knr1997/quiz-tracker-apiserver/global"
)

type SysQuiz struct {
	global.GVA_MODEL
	Name string    `json:"name" form:"name" gorm:"comment:Quiz Name"`
	Date time.Time `json:"date"`

	// Foreign Key
	CourseID uint      `json:"courseId" gorm:"not null;comment:Course ID"`
	Course   SysCourse `json:"course" gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

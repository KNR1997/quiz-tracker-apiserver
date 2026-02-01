package system

import (
	"github.com/knr1997/quiz-tracker-apiserver/database"
	"github.com/knr1997/quiz-tracker-apiserver/model/common/request"
	"github.com/knr1997/quiz-tracker-apiserver/model/system"
)

type CourseService struct{}

var CourseServiceApp = new(CourseService)

func (courseService *CourseService) CreateCourse(e system.SysCourse) (err error) {
	err = database.DB.Create(&e).Error
	return err
}

func (courseService *CourseService) UpdateCourse(e *system.SysCourse) (err error) {
	err = database.DB.Save(e).Error
	return err
}

func (courseService *CourseService) GetCourse(id uint) (course system.SysCourse, err error) {
	err = database.DB.Where("id = ?", id).First(&course).Error
	return
}

func (courseService *CourseService) GetCourseInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := database.DB.Model(&system.SysCourse{})
	var courseList []system.SysCourse

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&courseList).Error
	return courseList, total, err
}

func (courseService *CourseService) DeleteCourse(e system.SysCourse) (err error) {
	err = database.DB.Delete(&e).Error
	return err
}

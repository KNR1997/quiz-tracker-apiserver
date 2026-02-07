package system

import (
	"github.com/knr1997/quiz-tracker-apiserver/database"
	"github.com/knr1997/quiz-tracker-apiserver/model/common/request"
	"github.com/knr1997/quiz-tracker-apiserver/model/system"
)

type QuizService struct{}

var QuizServiceApp = new(QuizService)

func (quizService *QuizService) CreateQuiz(e system.SysQuiz) (err error) {
	err = database.DB.Create(&e).Error
	return err
}

func (quizService *QuizService) UpdateQuiz(e *system.SysQuiz) (err error) {
	err = database.DB.Save(e).Error
	return err
}

func (quizService *QuizService) GetQuizInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := database.DB.Model(&system.SysQuiz{})
	var quizList []system.SysQuiz

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Course").Find(&quizList).Error
	return quizList, total, err
}

func (quizService *QuizService) DeleteQuiz(id uint) error {
	return database.DB.Delete(&system.SysQuiz{}, id).Error
}

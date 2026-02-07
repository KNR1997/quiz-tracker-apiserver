package system

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/knr1997/quiz-tracker-apiserver/global"
	"github.com/knr1997/quiz-tracker-apiserver/model/common/request"
	"github.com/knr1997/quiz-tracker-apiserver/model/common/response"
	"github.com/knr1997/quiz-tracker-apiserver/model/system"
	"github.com/knr1997/quiz-tracker-apiserver/utils"
	"go.uber.org/zap"
)

type QuizApi struct{}

func (e *QuizApi) CreateQuiz(c *gin.Context) {
	var quiz system.SysQuiz
	err := c.ShouldBindJSON(&quiz)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// err = utils.Verify(course, utils.CustomerVerify)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	err = quizService.CreateQuiz(quiz)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Failed to create.", c)
		return
	}
	response.OkWithMessage("Created successfully.", c)
}

func (e *QuizApi) UpdateQuiz(c *gin.Context) {
	var input system.SysQuiz

	// Get ID from URL
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("Quiz ID is required", c)
		return
	}

	// Bind request body
	if err := c.ShouldBindJSON(&input); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// Assign ID manually
	input.ID = utils.StrToUint(id) // or uint conversion

	// Update
	if err := quizService.UpdateQuiz(&input); err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Failed to update.", c)
		return
	}

	response.OkWithMessage("Updated successfully.", c)
}

func (e *QuizApi) GetQuizList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// err = utils.Verify(pageInfo, utils.PageInfoVerify)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	quizList, total, err := quizService.GetQuizInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve."+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     quizList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Retrieved successfully.", c)
}

func (e *QuizApi) DeleteQuiz(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("Quiz ID is required", c)
		return
	}

	courseID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.FailWithMessage("Invalid course ID", c)
		return
	}

	if err := quizService.DeleteQuiz(uint(courseID)); err != nil {
		global.GVA_LOG.Error("Delete failed!", zap.Error(err))
		response.FailWithMessage("Failed to delete.", c)
		return
	}

	response.OkWithMessage("Deleted successfully.", c)
}

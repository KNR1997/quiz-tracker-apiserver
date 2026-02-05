package system

import (
	"github.com/gin-gonic/gin"
	"github.com/knr1997/quiz-tracker-apiserver/global"
	"github.com/knr1997/quiz-tracker-apiserver/model/common/request"
	"github.com/knr1997/quiz-tracker-apiserver/model/common/response"
	"github.com/knr1997/quiz-tracker-apiserver/model/system"
	sysRes "github.com/knr1997/quiz-tracker-apiserver/model/system/response"
	"github.com/knr1997/quiz-tracker-apiserver/utils"
	"go.uber.org/zap"
)

type CourseApi struct{}

func (e *CourseApi) CreateCourse(c *gin.Context) {
	var course system.SysCourse
	err := c.ShouldBindJSON(&course)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// err = utils.Verify(course, utils.CustomerVerify)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	err = courseService.CreateCourse(course)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Failed to create.", c)
		return
	}
	response.OkWithMessage("Created successfully.", c)
}

func (e *CourseApi) UpdateCourse(c *gin.Context) {
	var input system.SysCourse

	// Get ID from URL
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("Course ID is required", c)
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
	if err := courseService.UpdateCourse(&input); err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Failed to update.", c)
		return
	}

	response.OkWithMessage("Updated successfully.", c)
}

func (e *CourseApi) GetCourse(c *gin.Context) {
	var course system.SysCourse
	err := c.ShouldBindQuery(&course)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// err = utils.Verify(customer.GVA_MODEL, utils.IdVerify)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	data, err := courseService.GetCourse(course.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve.", c)
		return
	}
	response.OkWithDetailed(sysRes.CourseResponse{Course: data}, "Retrieved successfully.", c)
}

func (e *CourseApi) GetCourserList(c *gin.Context) {
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
	customerList, total, err := courseService.GetCourseInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve."+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     customerList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Retrieved successfully.", c)
}

func (e *CourseApi) DeleteCourse(c *gin.Context) {
	var course system.SysCourse
	err := c.ShouldBindJSON(&course)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// err = utils.Verify(course.GVA_MODEL, utils.IdVerify)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	err = courseService.DeleteCourse(course)
	if err != nil {
		global.GVA_LOG.Error("Failed to delete!", zap.Error(err))
		response.FailWithMessage("Failed to delete.", c)
		return
	}
	response.OkWithMessage("Deleted successfully.", c)
}

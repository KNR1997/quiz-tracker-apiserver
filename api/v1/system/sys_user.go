package system

import (
	"github.com/gin-gonic/gin"
	"github.com/knr1997/quiz-tracker-apiserver/global"
	"github.com/knr1997/quiz-tracker-apiserver/model/common/response"
	systemReq "github.com/knr1997/quiz-tracker-apiserver/model/system/request"
	"github.com/knr1997/quiz-tracker-apiserver/utils"
	"go.uber.org/zap"
)

type BaseApi struct{}

func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo systemReq.GetUserList
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userService.GetUserInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve.!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Successfully obtained", c)
}

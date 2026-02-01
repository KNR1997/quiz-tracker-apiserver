package example

import (
	"github.com/gin-gonic/gin"
	"github.com/knr1997/quiz-tracker-apiserver/global"
	"github.com/knr1997/quiz-tracker-apiserver/model/common/response"
	"github.com/knr1997/quiz-tracker-apiserver/model/example"
	"github.com/knr1997/quiz-tracker-apiserver/utils"
	"go.uber.org/zap"
)

type CustomerApi struct{}

// CreateExaCustomer
// @Tags      ExaCustomer
// @Summary   创建客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户用户名, 客户手机号码"
// @Success   200   {object}  response.Response{msg=string}  "创建客户"
// @Router    /customer/customer [post]
func (e *CustomerApi) CreateExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(customer, utils.CustomerVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customer.SysUserID = 1
	customer.SysUserAuthorityID = 1
	err = customerService.CreateExaCustomer(customer)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Failed to create.", c)
		return
	}
	response.OkWithMessage("Created successfully.", c)
}

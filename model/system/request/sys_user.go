package request

import (
	common "github.com/knr1997/quiz-tracker-apiserver/model/common/request"
)

type GetUserList struct {
	common.PageInfo
	// Username string `json:"username" form:"username"`
	// NickName string `json:"nickName" form:"nickName"`
	// Phone    string `json:"phone" form:"phone"`
	// Email    string `json:"email" form:"email"`
}

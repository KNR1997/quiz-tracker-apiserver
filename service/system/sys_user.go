package system

import (
	"github.com/knr1997/quiz-tracker-apiserver/database"
	"github.com/knr1997/quiz-tracker-apiserver/model/system"
	systemReq "github.com/knr1997/quiz-tracker-apiserver/model/system/request"
)

type UserService struct{}

var UserServiceApp = new(UserService)

func (userService *UserService) GetUserInfoList(info systemReq.GetUserList) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := database.DB.Model(&system.SysUser{})
	var userList []system.SysUser

	// if info.NickName != "" {
	// 	db = db.Where("nick_name LIKE ?", "%"+info.NickName+"%")
	// }
	// if info.Phone != "" {
	// 	db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	// }
	// if info.Username != "" {
	// 	db = db.Where("username LIKE ?", "%"+info.Username+"%")
	// }
	// if info.Email != "" {
	// 	db = db.Where("email LIKE ?", "%"+info.Email+"%")
	// }

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return userList, total, err
}

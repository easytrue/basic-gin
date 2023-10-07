package dto

import (
	"basicGin/model"
)

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required"`
}

type UserAddDTO struct {
	ID       uint
	Name     string `json:"name" form:"name" binding:"required" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `json:"password,omitempty" form:"password" binding:"required"`
	NickName string `json:"nickname" form:"nickname" binding:"required"`
	Status   int    `json:"status" form:"status" binding:"required"`
	Avatar   string `json:"avatar,omitempty" form:"avatar"`
	Mobile   string `json:"mobile" form:"mobile" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
}

func (m *UserAddDTO) ConvertToModel(iUser *model.User) {
	iUser.Name = m.Name
	iUser.Password = m.Password
	iUser.NickName = m.NickName
	iUser.Status = m.Status
	iUser.Avatar = m.Avatar
	iUser.Mobile = m.Mobile
	iUser.Email = m.Email

}

// 用户列表 dto
type UserListDTO struct {
	Paginate
}

type UserUpdateDTO struct {
	ID       uint   `json:"id" form:"id" uri:"id"`
	Name     string `json:"name" form:"name"`
	NickName string `json:"nickname" form:"nickname"`
	Status   int    `json:"status" form:"status"`
	Avatar   string `json:"avatar" form:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
}

func (m *UserUpdateDTO) ConvertToModel(iUser *model.User) {
	iUser.ID = m.ID
	iUser.Name = m.Name
	iUser.NickName = m.NickName
	iUser.Status = m.Status
	iUser.Avatar = m.Avatar
	iUser.Mobile = m.Mobile
	iUser.Email = m.Email
}

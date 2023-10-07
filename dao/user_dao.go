package dao

import (
	"basicGin/dto"
	"basicGin/model"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{NewBaseDao()}
	}
	return userDao
}

func (m *UserDao) GetUserByName(stName string) (model.User, error) {
	var iUser model.User
	err := m.Orm.Model(&iUser).Where("name=?", stName).Find(&iUser).Error
	return iUser, err
}

func (m *UserDao) GetUserByNameAndPassword(stUserName, stUserPassword string) model.User {
	var iUser model.User
	m.Orm.Model(&iUser).Where("name=? and password=?", stUserName, stUserPassword).Find(&iUser)
	return iUser
}

func (m *UserDao) CheckUserNameExist(stUsername string) bool {
	var total int64
	m.Orm.Model(&model.User{}).Where("name=?", stUsername).Count(&total)
	return total > 0
}

func (m *UserDao) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	var iUser model.User
	iUserAddDTO.ConvertToModel(&iUser)
	err := m.Orm.Save(&iUser).Error
	if err == nil {
		iUserAddDTO.ID = iUser.ID
		iUserAddDTO.Password = ""
	}
	return err
}

// 查询详情
func (m *UserDao) GetUserById(id uint) (model.User, error) {
	var iUser model.User

	err := m.Orm.First(&iUser, id).Error

	return iUser, err

}

func (m *UserDao) GetUserList(iUserListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	var getUserList []model.User
	var total int64

	err := m.Orm.Model(&model.User{}).
		Scopes(Paginate(iUserListDTO.Paginate)).
		Find(&getUserList).
		Offset(-1).Limit(-1).
		Count(&total).Error

	return getUserList, total, err
}

func (m *UserDao) UpdateUser(iUserUpdateDTO *dto.UserUpdateDTO) error {
	var iUser model.User
	m.Orm.First(&iUser, iUserUpdateDTO.ID)
	iUserUpdateDTO.ConvertToModel(&iUser)
	return m.Orm.Save(&iUser).Error
}

func (m *UserDao) DeleteUserById(id uint) error {
	return m.Orm.Delete(&model.User{}, id).Error
}

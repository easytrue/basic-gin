package service

import (
	"basicGin/dao"
	"basicGin/dto"
	"basicGin/global"
	"basicGin/global/constans"
	"basicGin/model"
	"basicGin/utils"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}

	return userService
}

func setLoginUserTokenToRedis(uid uint, token string) error {
	return global.RedisClient.Set(
		strings.Replace(constans.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", strconv.Itoa(int(uid)), -1),
		token,
		viper.GetDuration("jwt.expiresAt")*time.Second,
	)
}

func (m *UserService) Login(iUserDTO dto.UserLoginDTO) (model.User, string, error) {
	var errResult error
	var token string
	iUser, err := m.Dao.GetUserByName(iUserDTO.Name)

	// 用户名或密码不正确
	if err != nil || !utils.CompareHashAndPassword(iUser.Password, iUserDTO.Password) {
		errResult = errors.New("account or password error")
	} else { // 登陆成功
		token, err = utils.GenerateToken(iUser.ID, iUser.Name)
		if err != nil {
			errResult = errors.New(fmt.Sprintf("Generate Token Error: %s", err.Error()))
		} else {
			err = setLoginUserTokenToRedis(iUser.ID, token)
			if err != nil {
				errResult = errors.New(fmt.Sprintf("Token REDIS SET ERROR: %s", err.Error()))
			}
		}
	}
	return iUser, token, errResult
}

func (m *UserService) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	if m.Dao.CheckUserNameExist(iUserAddDTO.Name) {
		return errors.New("User Name Exist")
	}
	return m.Dao.AddUser(iUserAddDTO)
}

func (m *UserService) GetUserById(iCommonIDDTO *dto.CommonIDDTO) (model.User, error) {
	return m.Dao.GetUserById(iCommonIDDTO.ID)
}

func (m *UserService) GetUserList(iUserListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	return m.Dao.GetUserList(iUserListDTO)
}

func (m *UserService) UpdateUser(iUserUpdateDTO *dto.UserUpdateDTO) error {
	if iUserUpdateDTO.ID == 0 {
		return errors.New("Invalid User ID")
	}
	return m.Dao.UpdateUser(iUserUpdateDTO)
}

func (m *UserService) DeleteUserById(iCommonIDDTO *dto.CommonIDDTO) error {
	return m.Dao.DeleteUserById(iCommonIDDTO.ID)
}

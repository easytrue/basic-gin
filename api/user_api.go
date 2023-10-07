package api

import (
	"basicGin/dto"
	"basicGin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ADD_USER_ERROR           = 10011
	NOT_FUND_USER_ERROR      = 10012
	NOT_FUND_LIST_USER_ERROR = 10013
	UPDATE_USER_ERROR        = 10014
	DELETE_USER_ERROR        = 10015
	LOGIN_USER_ERROR         = 10016
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

// Login
// @Summary 用户登陆
// @Description 用户登陆接口
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登陆成功"
// @Failure 401 {string} string "登陆失败"
// @Router /api/v1/public/user/login [post]
func (m UserApi) Login(c *gin.Context) {
	var isUserLoginDTO dto.UserLoginDTO

	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &isUserLoginDTO}).GetError; err == nil {
		fmt.Println(err)
		return
	}

	iUser, token, err := m.Service.Login(isUserLoginDTO)

	if err != nil {
		m.Fail(ResponseJson{
			Status: http.StatusUnauthorized,
			Code:   LOGIN_USER_ERROR,
			Msg:    err.Error(),
		})
		return
	}

	m.Success(ResponseJson{
		Code: 200,
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
	})
}

func (m UserApi) AddUser(c *gin.Context) {
	var iUserAddDTO dto.UserAddDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserAddDTO}).GetError(); err != nil {
		return
	}

	//file, _ := c.FormFile("file")
	//stFilePath := fmt.Sprintf("./upload/%s", file.Filename)
	//_ = c.SaveUploadedFile(file, stFilePath)
	//iUserAddDTO.Avatar = stFilePath

	err := m.Service.AddUser(&iUserAddDTO)
	if err != nil {
		m.ServiceFail(ResponseJson{
			Code: ADD_USER_ERROR,
			Msg:  err.Error(),
		})
		return
	}
	m.Success(ResponseJson{
		Data: iUserAddDTO,
	})
}

func (m *UserApi) GetUserById(c *gin.Context) {

	fmt.Println(c.Get("auth_user"))

	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}
	iUser, err := m.Service.GetUserById(&iCommonIDDTO)
	if err != nil {
		m.ServiceFail(ResponseJson{
			Code: NOT_FUND_USER_ERROR,
			Msg:  "没有用户信息",
		})
		return
	}

	m.Success(ResponseJson{
		Data: iUser,
	})

}

func (m *UserApi) GetUserList(c *gin.Context) {
	var iUserListDto dto.UserListDTO

	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserListDto}).GetError(); err != nil {
		return
	}

	getUserList, total, err := m.Service.GetUserList(&iUserListDto)

	if err != nil {
		m.ServiceFail(ResponseJson{
			Code: NOT_FUND_LIST_USER_ERROR,
			Msg:  "没有用户信息",
		})
		return
	}

	m.Success(ResponseJson{
		Data:  getUserList,
		Total: total,
	})
}

func (m UserApi) UpdateUser(c *gin.Context) {
	var iUserUpdateDTO dto.UserUpdateDTO

	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserUpdateDTO, BindAll: true}).GetError(); err != nil {
		return
	}

	fmt.Println(iUserUpdateDTO.ID, iUserUpdateDTO)

	err := m.Service.UpdateUser(&iUserUpdateDTO)
	if err != nil {
		m.ServiceFail(ResponseJson{
			Code: UPDATE_USER_ERROR,
			Msg:  err.Error(),
		})
		return
	}
	m.Success(ResponseJson{
		Data: iUserUpdateDTO,
	})
}

func (m UserApi) DeleteUserById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO

	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	err := m.Service.DeleteUserById(&iCommonIDDTO)
	if err != nil {
		m.ServiceFail(ResponseJson{
			Code: DELETE_USER_ERROR,
			Msg:  err.Error(),
		})
		return
	}
	m.Success(ResponseJson{
		Data: "",
	})
}

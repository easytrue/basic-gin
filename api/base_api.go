package api

import (
	"basicGin/global"
	"basicGin/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"reflect"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

type BuildRequestOption struct {
	Ctx     *gin.Context
	DTO     any
	BindUri bool
	BindAll bool
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: global.Logger,
	}
}

func (m *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {
	var errResult error
	// 绑定请求上下文
	m.Ctx = option.Ctx
	// 绑定请求数据
	if option.DTO != nil {
		if option.BindUri || option.BindAll {
			errResult = utils.AppendError(errResult, m.Ctx.ShouldBindUri(option.DTO))
		}
		if option.BindAll || !option.BindUri {
			errResult = utils.AppendError(errResult, m.Ctx.ShouldBind(option.DTO))
		}
	}

	if errResult != nil {
		errResult = m.parseValidateErrors(errResult, option.DTO)
		m.AddError(errResult)
		m.Fail(ResponseJson{
			Msg: m.GetError().Error(),
		})
	}

	return m
}

func (m *BaseApi) AddError(errNew error) {
	m.Errors = utils.AppendError(m.Errors, errNew)
}

func (m *BaseApi) GetError() error {
	return m.Errors
}

func (m *BaseApi) parseValidateErrors(errs error, target any) error {
	var errResult error
	errValidation, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errs
	}
	// 通过反射获取指针指向元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errValidation {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}
		if errMessage == "" {
			errMessage = fmt.Sprintf("%s:%s Error", fieldErr.Field(), fieldErr.Tag())
		}
		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}

	return errResult
}

func (m *BaseApi) Fail(response ResponseJson) {
	Fail(m.Ctx, response)
}

func (m *BaseApi) Success(response ResponseJson) {
	Success(m.Ctx, response)
}

func (m *BaseApi) ServiceFail(response ResponseJson) {
	ServiceFail(m.Ctx, response)
}

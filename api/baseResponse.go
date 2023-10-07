package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type ResponseJson struct {
	Status int    `json:"-"`
	Code   int    `json:"code"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
	Total  int64  `json:"total,omitempty""`
}

func (m ResponseJson) IsEmpty() bool {
	return reflect.DeepEqual(m, ResponseJson{})
}

func Success(context *gin.Context, responseJson ResponseJson) {
	HttpResponse(context, getStatus(responseJson, http.StatusOK), responseJson)
}

func ServiceFail(context *gin.Context, responseJson ResponseJson) {
	HttpResponse(context, getStatus(responseJson, http.StatusInternalServerError), responseJson)
}

func Fail(context *gin.Context, responseJson ResponseJson) {
	HttpResponse(context, getStatus(responseJson, http.StatusBadRequest), responseJson)
}

func getStatus(responseJson ResponseJson, defaultStatus int) int {
	if responseJson.Status == 0 {
		return defaultStatus
	}
	return responseJson.Status
}

func HttpResponse(context *gin.Context, status int, responseJson ResponseJson) {
	if responseJson.IsEmpty() {
		context.AbortWithStatus(status)
		return
	}
	context.AbortWithStatusJSON(status, responseJson)
}

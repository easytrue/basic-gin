package middlewares

import (
	"basicGin/api"
	"basicGin/global"
	"basicGin/global/constans"
	"basicGin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取 token
		tokenString := context.GetHeader("Authorization")
		if strings.HasPrefix(tokenString, "Bearer") {
			tokenString = strings.Split(tokenString, " ")[1]
		}

		claims, err := utils.ParseToken(tokenString)
		userId := claims.ID
		if err != nil || userId == 0 {
			authError(context, err)
			// context.Abort()
			return
		}
		// token 是否过期 是否相同
		stUserId := strconv.Itoa(int(userId))
		redisTokenKey := strings.Replace(constans.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", stUserId, -1)
		redisToken, err := global.RedisClient.Get(redisTokenKey)

		if redisToken != tokenString || err != nil {
			authError(context, err)
			// context.Abort()
			return
		}
		context.Set("auth_users", *claims)
		context.Set("auth_user_id", userId)
		context.Next()

	}
}

func authError(context *gin.Context, err error) {
	api.Fail(context, api.ResponseJson{
		Status: http.StatusUnauthorized,
		Code:   http.StatusUnauthorized,
		Msg:    err.Error(),
	})
}

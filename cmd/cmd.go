package cmd

import (
	"basicGin/conf"
	"basicGin/global"
	"basicGin/router"
	"basicGin/utils"
	"fmt"
)

func Start() {
	var initError error

	// 初始化配置文件
	conf.InitConfig()
	// 初始化日志组件
	global.Logger = conf.InitLogger()
	// 数据库链接
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initError = utils.AppendError(initError, err)
	}
	// 初始化 redis
	redisClient, err := conf.InitRedis()
	global.RedisClient = redisClient
	if err != nil {
		initError = utils.AppendError(initError, err)
	}

	// 判断初始化错误
	if initError != nil {
		if global.Logger != nil {
			global.Logger.Error(initError.Error())
		}
		panic(initError.Error())
	}

	// 初始化系统路由
	router.InitRouter()
}

func Clear() {
	fmt.Println("====clear====")
}

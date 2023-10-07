package router

import (
	_ "basicGin/docs"
	"basicGin/global"
	middlewares "basicGin/middleware"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type IFnRegisterRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRoutes []IFnRegisterRoute
)

// 初始化系统路由
func InitRouter() {
	// 创建监听 ctrl + c, 应用退出信号上下文
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	// 初始化 gin 框架,并主持相关路由
	r := gin.Default()
	// 跨域支持
	r.Use(middlewares.Cors())
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")
	rgAuth.Use(middlewares.JWTAuth())

	// 初始化基础路由
	initBasePlatformRoutes()
	// 注册自定义验证器
	cunstomValidator()

	// 遍历路由结构体,开始注册系统各模块对应的路由信息
	for _, fRegisterRoute := range gfnRoutes {
		fRegisterRoute(rgPublic, rgAuth)
	}
	// 集成swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 读取配置 端口
	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8999"
	}
	// 创建 web server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	// 开启协程启动 goroutine 来开启 web 服务,避免主线程的信号监听被阻塞
	go func() {
		global.Logger.Info(fmt.Sprintf("Start Listen: %s", stPort))
		// server 监听
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// TODO: 记录日志
			global.Logger.Error(fmt.Sprintf("Start Server Error: %s", err.Error()))
			return
		}
	}()
	// 空结构体
	<-ctx.Done()
	// cancelCtx()
	// 停止 server
	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(ctx); err != nil {
		// TODO: 记录日志
		global.Logger.Error(fmt.Sprintf("Stop Server Error: %s", err.Error()))
		return
	}
	global.Logger.Info("Stop Server Success")
}

// 注册路由回调
func RegisterRoute(fn IFnRegisterRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

func initBasePlatformRoutes() {

	InitUserRouter()
}

// 自定义验证器示例
func cunstomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("first_is_a", func(fl validator.FieldLevel) bool {
			if value, ok := fl.Field().Interface().(string); ok {
				if value != "" && 0 == strings.Index(value, "a") {
					return true
				}
			}
			return false
		})
	}
}

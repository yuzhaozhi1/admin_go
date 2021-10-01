package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/middleware"
	"net/http"
)

// router 路由

// Routers gin 初始化路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	// 静态的文件代理, 可以用 nginx 来代理, 为用户头像和文件提供静态地址
	Router.StaticFS(global.GLOBAL_CONFIG.Local.Path, http.Dir(global.GLOBAL_CONFIG.Local.Path))
	// 使用 https
	// Router.Use(middleware.LoadTls())
	// global.GLOBAL_LOG.Info("open https middleware ok! ")

	// 跨域
	Router.Use(middleware.Cors())
	global.GLOBAL_LOG.Info("use cors middleware ok!")

	// 文档地址
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GLOBAL_LOG.Info("register swagger handler")

	return Router
}

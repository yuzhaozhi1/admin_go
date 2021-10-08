package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yuzhaozhi1/admin_go/api/v1"
)

// Base 路由

// InitBaseRouter base 路由
 func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	 BaseRouter := Router.Group("base")
	 {
		 BaseRouter.GET("/captcha", v1.Captcha)  // 生成验证码
		 BaseRouter.POST("/login", v1.Login)  // 用户登录
		 BaseRouter.POST("register", v1.Register)
	 }
	 return BaseRouter
 }

package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yuzhaozhi1/admin_go/api/v1"
)

// 用户相关的高级操作路由

func InitUserRouter(Router *gin.RouterGroup)  {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("changePassword", v1.ChangePassword) // 用户修改密码
	}

}

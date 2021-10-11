package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/model/request"
	"github.com/yuzhaozhi1/admin_go/model/response"
	"github.com/yuzhaozhi1/admin_go/service"
)

// casbin 拦截器

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var claims interface{}
		var ok bool
		if claims, ok = c.Get("claims"); !ok {
			c.Abort()
		}
		waitUse := claims.(*request.CustomClaims)
		// 获取请求的url
		obj := c.Request.URL.RequestURI()

		// 获取请求的方法
		act := c.Request.Method

		// 获取用户的角色
		sub := waitUse.AuthorityId

		e := service.Casbin()

		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.GLOBAL_CONFIG.System.Env == "develop" || success {
			c.Next()
		}else {
			response.FailWithDetailed(gin.H{}, "权限不足",c)
			c.Abort()
		}
	}
}

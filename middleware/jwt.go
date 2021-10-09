package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yuzhaozhi1/admin_go/model/response"
)

// JWTAuth  jwt 用户校验中间件
func JWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 从jwt 鉴权取 头部token 数据, key为 (Authorization:不用), x-token, 登录时返回token信息
		token := c.Request.Header.Get("x-token")

		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "请求不合法! 未携带token", c)
			return
		}
	}

}

type JWT struct {
	SigningKey []byte
}

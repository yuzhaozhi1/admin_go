package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/model/request"
	"github.com/yuzhaozhi1/admin_go/model/response"
	"github.com/yuzhaozhi1/admin_go/service"
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

		// 校验token是否在黑名单中,
		if service.IsInBlackList(token) {
			global.GLOBAL_LOG.Error("黑名单用户访问了" + c.Request.Host)
			response.FailWithDetailed(gin.H{"reload": true}, "您的账户异地登录或令牌失效", c)
			c.Abort() // 防止后面的函数调用
		}

	}

}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpiredErr     = errors.New("token is expired")           // token 过期错误
	TokenNotValidYetErr = errors.New("token not active yet")       // 令牌尚未生效
	TokenMalformedErr   = errors.New("that's not even a token")    // 令牌格式错误
	TokenInvalidErr     = errors.New("couldn't handel this token") // 令牌无效
)

// NewJwt 返回一个Jwt 对象
func NewJwt() *JWT {
	return &JWT{
		SigningKey: []byte(global.GLOBAL_CONFIG.Jwt.SigningKey),
	}
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 用旧的token来换取新的token, 利用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.GLOBA_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析token,将token 反解为结构体
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	// 解析err 错误的类型,进行返回
	if err != nil {
		// err 断言类型
		if validationError, ok := err.(*jwt.ValidationError); ok {
			// 参与运算的两数各对应的二进位相与。(两位均为1才为1）
			// jwt.ValidationErrorMalformed = 1
			// 就是抛出的错误可以是符合的错误  err1|err2   这时使用 & 表示如果 抛出的错误是其中的一个就可以
			if validationError.Errors & jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformedErr
			} else if validationError.Errors & jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpiredErr
			} else if validationError.Errors & jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYetErr
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid{
			return claims, nil
		}
	}
	return nil, TokenInvalidErr
}

/*
	// 0  1, 2,  3,   4,  5,   6     7    8     9    10     11    12   13
	// 0, 1, 10, 11, 100, 101, 110, 111, 1000, 1001, 1010, 1011, 1100, 1101
	// err1 和 err2  是一个类型的错误
	// err1= 5  =>  101
	// err2= 3  =>  011
	// new_err  => 5|3 =>  101|011  => 111 => 7
	// 这时我们来了一个新的错误是 7 如果直接判断 new_err == err1 肯定是不成立的
	// new_err & err1  => 7 & 5 =>  111 & 101  => 101 != 0  所以 new_err 可以看成是属于 err1 和err3 一个类型的
*/

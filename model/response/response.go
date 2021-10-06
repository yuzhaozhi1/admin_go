package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 统一返回数据的格式

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 0
	ERR0R   = 7
)

// Result 返回数据
func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func Fail(c *gin.Context) {
	Result(ERR0R, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(ERR0R, map[string]interface{}{}, msg, c)
}

func FailWithDetailed(data interface{}, msg string, c *gin.Context) {
	Result(ERR0R, data, msg, c)
}

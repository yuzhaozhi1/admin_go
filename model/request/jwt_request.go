package request

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// Jwt request

// CustomClaims 自定义声明
type CustomClaims struct {
	UUID     uuid.UUID
	ID       uint
	Username string
	NickName string
	AuthorityId string
	BufferTime int64
	jwt.StandardClaims   // jwt 标准声明
}

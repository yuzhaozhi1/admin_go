package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

type server interface {
	ListenAndServe() error
}

func InitServer(address string, c *gin.Engine) server {
	s := endless.NewServer(address, c)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20 // 1024 * 1024
	return s
}

// func ()  {
//
// }

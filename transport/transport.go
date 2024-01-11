package transport

import "github.com/gin-gonic/gin"

type TransportHandle func(gin.Context) (interface{}, error)

type Transport struct {
	engine *gin.Engine
}

func New(e *gin.Engine) *Transport {
	return &Transport{
		engine: e,
	}
}

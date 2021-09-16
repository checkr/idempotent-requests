package http

import "github.com/gin-gonic/gin"

type PingHandler struct{}

func NewPingHandler() Handler {
	return PingHandler{}
}

func (PingHandler) Handle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

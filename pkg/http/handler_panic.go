package http

import "github.com/gin-gonic/gin"

type PanicHandler struct{}

func NewPanicHandler() Handler {
	return PanicHandler{}
}

func (PanicHandler) Handle(c *gin.Context) {
	panic("Oh no!")
}

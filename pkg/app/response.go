package app

import (
	"github.com/gin-gonic/gin"
)

// Gin :
type Gin struct {
	C *gin.Context
}

// Response :
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response : send response after request
func (g *Gin) Response(httpCode int, errMsg string, data interface{}) interface{} {
	response := Response{
		Code: httpCode,
		Msg:  errMsg,
		Data: data,
	}
	g.C.JSON(httpCode, response)
	return response
}

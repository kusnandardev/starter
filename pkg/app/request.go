package app

import (
	"fmt"
	"kusnandartoni/starter/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// MarkErrors :
func MarkErrors(errors []*validation.Error) string {
	res := ""
	for _, err := range errors {
		res = fmt.Sprintf("%s %s", err.Key, err.Message)
	}

	return res
}

// GetClaims :
func GetClaims(c *gin.Context) util.Claims {
	var clm util.Claims
	claims, _ := c.Get("claims")
	mapstructure.Decode(claims, &clm)
	return clm
}

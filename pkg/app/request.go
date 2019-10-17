package app

import (
	"kusnandartoni/starter/pkg/util"
	"fmt"

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
	claims, exist := c.Get("claims")

	mapstructure.Decode(claims, &clm)
	clm.UUID = clm.Id
	if exist {
		fmt.Println(clm)
	}
	return clm
}

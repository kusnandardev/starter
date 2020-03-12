package util

import (
	"kusnandartoni/starter/pkg/setting"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

// GetPage :
func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize

	}
	return result
}

// GetPerPage :
func GetPerPage(c *gin.Context) int {
	result := 0
	perPage := com.StrTo(c.Query("perpage")).MustInt()
	if perPage > 0 {
		result = perPage
	} else {
		result = 25
	}
	return result
}

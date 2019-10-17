package app

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid :
func BindAndValid(c *gin.Context, form interface{}) (int, string) {
	err := c.Bind(form)

	if err != nil {
		return http.StatusBadRequest, fmt.Sprintf("invalid request param: %v", err)
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, fmt.Sprintf("internal server error: %v", err)
	}
	if !check {
		return http.StatusBadRequest, MarkErrors(valid.Errors)
	}
	return http.StatusOK, "ok"
}

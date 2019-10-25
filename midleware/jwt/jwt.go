package jwt

import (
	"kusnandartoni/starter/pkg/setting"
	"kusnandartoni/starter/pkg/util"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT :
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = http.StatusOK
		msg := ""
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = http.StatusNetworkAuthenticationRequired
			msg = "Auth Token Required"
		} else {
			claims, err := util.ParseToken(token)
			code = http.StatusUnauthorized
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					msg = "Token Expired"
				default:
					msg = "Token Failed"
				}
			} else {
				valid := claims.VerifyIssuer(setting.AppSetting.Issuer, true)
				if !valid {
					msg = "Issuer is not valid"
				}
				c.Set("claims", claims)
			}
		}

		if code != http.StatusOK {
			c.JSON(code, gin.H{
				"code": code,
				"msg":  msg,
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

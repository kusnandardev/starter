package jwt

import (
	"kusnandartoni/starter/pkg/logging"
	"kusnandartoni/starter/pkg/setting"
	"kusnandartoni/starter/pkg/util"
	"kusnandartoni/starter/redisdb"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT :
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			logger        = logging.Logger{UUID: "SYS"}
			data          interface{}
			code          = http.StatusOK
			msg           = ""
			token         = ""
			authorization = c.Request.Header.Get("Authorization")
		)

		splitedAuthorization := strings.Split(authorization, " ")
		if len(splitedAuthorization) != 2 {
			code = http.StatusUnauthorized
			msg = "Yout Auth Token is Invalid"
		}
		if splitedAuthorization[0] != "Bearer" {
			logger.Error("Your basic auth is invalid")
			return
		}
		token = splitedAuthorization[1]

		data = map[string]string{
			"token": token,
		}

		if token == "" {
			code = http.StatusUnauthorized
			msg = "Auth Token Required"
		} else {
			existToken := redisdb.GetSession(token)
			if existToken == "" {
				code = http.StatusUnauthorized
				msg = "Token Failed"
			}
			claims, err := util.ParseToken(token)
			if err != nil {
				code = http.StatusUnauthorized
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					msg = "Token Expired"
				default:
					msg = "Token Failed"
				}
			} else {
				valid := claims.VerifyIssuer(setting.AppSetting.Issuer, true)
				if !valid {
					code = http.StatusUnauthorized
					msg = "Issuer is not valid"
				}
				c.Set("claims", claims)
			}
		}

		if code != http.StatusOK {
			resp := gin.H{
				"code": code,
				"msg":  msg,
				"data": data,
			}
			c.JSON(code, resp)

			logger.Error(util.Stringify(resp))
			c.Abort()
			return
		}

		c.Next()
	}
}

package util

import (
	"fmt"
	"kusnandartoni/starter/pkg/setting"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

// Claims :
type Claims struct {
	jwt.StandardClaims
	UUID string `json:"uuid,omitempty"`
}

// GenerateToken :
func GenerateToken(id int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{}
	claims.Id = strconv.Itoa(int(id))
	claims.ExpiresAt = expireTime.Unix()
	claims.Issuer = setting.AppSetting.Issuer

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken :
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// GetEmailToken :
func GetEmailToken(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}

// ParseEmailToken :
func ParseEmailToken(token string) string {
	tkn, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	claims, _ := tkn.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%s", claims["email"])
}

package CONF

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 定义JWT过期时间
const JWTExpirationTime = 24 * time.Hour

type User struct {
	Uid      uint
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成jwt
func GenerateJWT(username string, userID uint) (string, error) {
	claims := User{
		userID,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(GetConfig().JWTSecret))
	return s, err
}

// 解析jwt
func ParseJWT(tokenString string) (*User, error) {
	t, err := jwt.ParseWithClaims(tokenString, &User{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetConfig().JWTSecret), nil
	})
	if claims, ok := t.Claims.(*User); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
)

type EasyToken struct {
	Username string
	Uid      int64
	Expires  int64
}

var (
	verifyKey  string
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
	ErrExpired = "token expired" // 令牌过期
	ErrOther   = "other error"   // 其他错误
)

func init() {
	// 这个是秘钥,可以配置到环境变量中
	verifyKey = "12345678"
}

// 获取token
func (e *EasyToken) GetToken() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: e.Expires, //time.Unix(c.ExpiresAt, 0)
		Issuer:    e.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(verifyKey))
	if err != nil {
		log.Println(err)
	}
	return tokenString, err
}

// 验证是否正确
func (e *EasyToken) ValidateToken(tokenString string) (bool, error) {
	if tokenString == "" {
		return false, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(verifyKey), nil
	})
	if token == nil {
		return false, errors.New(ErrInvalid)
	}
	if token.Valid {

		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New(ErrInvalid)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New(ErrExpired)
		} else {
			return false, errors.New(ErrOther)
		}
	} else {
		return false, errors.New(ErrOther)
	}
}

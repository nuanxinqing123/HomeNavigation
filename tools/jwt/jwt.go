// -*- coding: utf-8 -*-
// @Time    : 2022/3/14 17:19
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : jwt.go

package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// TokenExpireDuration Token过期时间(30天)
const TokenExpireDuration = time.Hour * 720

// 加盐
var mySecret = []byte("XOJogCiP1SJDxGXyHB")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(pwd string) (string, error) {
	// 创建声明数据
	c := MyClaims{
		Password: pwd,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "nuanxinqing",                              // 签发人
		},
	}

	// 使用指定的签名方式创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 使用指定的签名并获得完整编码后的Token
	return token.SignedString(mySecret)
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析Token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}

	// 校验Token
	if token.Valid {
		return mc, nil
	}

	return nil, errors.New("invalid token")
}

// -*- coding: utf-8 -*-
// @Time    : 2022/3/14 17:22
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : auth.go

package middleware

import (
	"Gin_HomeNavigation/tools/jwt"
	res "Gin_HomeNavigation/tools/response"
	"strings"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "Password"

// UserAuth 基于JWT的认证中间件
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			res.ResError(c, res.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			res.ResError(c, res.CodeInvalidToken)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			res.ResError(c, res.CodeInvalidToken)
			c.Abort()
			return
		}

		// 将当前请求的Password信息保存到请求的上下文c上
		c.Set(CtxUserIDKey, mc.Password)
		c.Next() // 后续的处理函数可以用过c.Get(CtxUserIDKey)来获取当前请求的用户信息
	}
}

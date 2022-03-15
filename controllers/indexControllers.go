// -*- coding: utf-8 -*-
// @Time    : 2021/11/15 17:16
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : indexControllers.go

package controllers

import (
	"Gin_HomeNavigation/dataSource"
	"Gin_HomeNavigation/tools/jwt"
	res "Gin_HomeNavigation/tools/response"
	val "Gin_HomeNavigation/tools/validator"
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Pwd struct {
	Password string `json:"password" binding:"required"`
}

// NoRouter 空路由
func NoRouter(ctx *gin.Context) {
	ctx.String(http.StatusOK, "The visit address does not exist, please return to the homepage")
}

// WebData 无密码数据
func WebData(c *gin.Context) {
	// 加载配置
	Web := dataSource.LoadConfig()

	if Web.SoftWare.Password == "" {
		res.ResSuccess(c, gin.H{
			"logo":       Web.Index.Logo,
			"title":      Web.Index.Title,
			"favicon":    Web.Index.Favicon,
			"mode":       Web.SoftWare.Mode,
			"web_data":   Web.Data,
			"text":       Web.Footer,
			"Background": Web.FooterStyle.Background,
			"LColor":     Web.FooterStyle.LColor,
			"SColor":     Web.FooterStyle.SColor,
			"FColor":     Web.FooterStyle.FColor,
			"IsLogin":    1,
		})
	} else {
		res.ResError(c, res.CodeNeedLogin)
	}
}

// PwdWebData 网站数据
func PwdWebData(c *gin.Context) {
	// 加载配置
	Web := dataSource.LoadConfig()

	res.ResSuccess(c, gin.H{
		"logo":       Web.Index.Logo,
		"title":      Web.Index.Title,
		"favicon":    Web.Index.Favicon,
		"mode":       Web.SoftWare.Mode,
		"web_data":   Web.Data,
		"text":       Web.Footer,
		"Background": Web.FooterStyle.Background,
		"LColor":     Web.FooterStyle.LColor,
		"SColor":     Web.FooterStyle.SColor,
		"FColor":     Web.FooterStyle.FColor,
		"IsLogin":    1,
	})
}

// LoginData 登录数据
func LoginData(c *gin.Context) {
	// 加载配置
	Web := dataSource.LoadConfig()

	res.ResSuccess(c, gin.H{
		"logo":       Web.Index.Logo,
		"title":      Web.Index.Title,
		"favicon":    Web.Index.Favicon,
		"text":       Web.Footer,
		"Background": Web.FooterStyle.Background,
		"LColor":     Web.FooterStyle.LColor,
		"SColor":     Web.FooterStyle.SColor,
		"FColor":     Web.FooterStyle.FColor,
	})
}

// Login 登录
func Login(c *gin.Context) {
	// 加载配置
	Web := dataSource.LoadConfig()

	// 获取参数
	p := new(Pwd)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 参数校验
		zap.L().Error("SignInHandle with invalid param", zap.Error(err))

		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.ResError(c, res.CodeInvalidParam)
			return
		}

		// 翻译错误
		res.ResErrorWithMsg(c, res.CodeInvalidParam, val.RemoveTopStruct(errs.Translate(val.Trans)))
		return
	}

	if Web.SoftWare.Password != p.Password {
		res.ResError(c, res.CodePassWordWrong)
		return
	} else {
		// 密码正确, 生成Token并返回
		token, err := jwt.GenToken(p.Password)
		if err != nil {
			zap.L().Error("Error: 生成JWT发生错误")
			res.ResError(c, res.CodeServerBusy)
			return
		}

		res.ResSuccess(c, gin.H{
			"Token": token,
		})
	}
}

// LoginF 登录路径
func LoginF(c *gin.Context) {
	// 加载配置
	Web := dataSource.LoadConfig()

	if Web.SoftWare.Password == "" {
		res.ResSuccess(c, gin.H{
			"IsLogin": 0,
		})
	} else {
		res.ResSuccess(c, gin.H{
			"IsLogin": 1,
		})
	}
}

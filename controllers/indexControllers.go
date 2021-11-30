// -*- coding: utf-8 -*-
// @Time    : 2021/11/15 17:16
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : indexControllers.go

package controllers

import (
	"Gin_HomeNavigation/dataSource"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Index 首页
func Index(ctx *gin.Context) {
	// 加载配置
	Web := dataSource.LoadConfig()

	// 判断是否启用密码登录
	if Web.SoftWare.Password == "" {
		// 渲染前端
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"logo":       Web.Index.Logo,
			"title":      Web.Index.Title,
			"favicon":    Web.Index.Favicon,
			"data":       Web.Data,
			"text":       Web.Footer,
			"Background": Web.FooterStyle.Background,
			"AColor":     Web.FooterStyle.AColor,
			"IsLogin":    1,
		})
	} else {
		// 启用密码登录，判断Session
		// 获取Session，判断用户是否登录
		isLogin := GetSession(ctx)

		if isLogin == true {
			// 已登录
			// 渲染前端
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"logo":       Web.Index.Logo,
				"title":      Web.Index.Title,
				"favicon":    Web.Index.Favicon,
				"data":       Web.Data,
				"text":       Web.Footer,
				"Background": Web.FooterStyle.Background,
				"AColor":     Web.FooterStyle.AColor,
				"IsLogin":    isLogin,
			})
		} else {
			// 未登录
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"logo":       Web.Index.Logo,
				"title":      Web.Index.Title,
				"favicon":    Web.Index.Favicon,
				"text":       Web.Footer,
				"Background": Web.FooterStyle.Background,
				"AColor":     Web.FooterStyle.AColor,
			})
		}
	}
}

// Login 登录
func Login(ctx *gin.Context) {
	// 加载配置
	Web := dataSource.LoadConfig()

	pwd := ctx.PostForm("password")
	if Web.SoftWare.Password != pwd {
		ctx.JSON(http.StatusOK, gin.H{
			// 密码错误
			"msg": "Wrong Password",
		})
	} else {
		session := sessions.Default(ctx)
		session.Set("userLogin", Web.SoftWare.Password)
		err := session.Save()
		if err != nil {
			log.Println(err)
			return
		}

		// 下放Session, 301跳转
		ctx.Redirect(http.StatusMovedPermanently, "/")
	}
}

// NoRouter 空路由
func NoRouter(ctx *gin.Context) {
	ctx.String(http.StatusOK, "The visit address does not exist, please return to the homepage")
}

// -*- coding: utf-8 -*-
// @Time    : 2021/11/15 17:16
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : indexControllers.go

package controllers

import (
	"Gin_HomeNavigation/dataSource"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	// 加载配置
	Web := dataSource.LoadConfig()

	// 渲染前端
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"logo":        Web.Index.Logo,
		"title":       Web.Index.Title,
		"favicon":     Web.Index.Favicon,
		"data":        Web.Data,
		"text":        Web.Footer,
		"Background":  Web.FooterStyle.Background,
		"AColor":      Web.FooterStyle.AColor,
		"AColorHover": Web.FooterStyle.AColorHover,
	})
}

func NoRouter(ctx *gin.Context) {
	ctx.String(http.StatusOK, "The visit address does not exist, please return to the homepage")
}

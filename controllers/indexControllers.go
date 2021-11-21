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
		"logo":                  Web.Config.Logo,
		"title":                 Web.Config.Title,
		"favicon":               Web.Config.Favicon,
		"FooterTitle":           Web.Config.FooterTitle,
		"FooterTitleLink":       Web.Config.FooterTitleLink,
		"FooterTextDataOne":     Web.Config.FooterTextDataOne,
		"FooterTextDataTwo":     Web.Config.FooterTextDataTwo,
		"FooterTextDataTwoLink": Web.Config.FooterTextDataTwoLink,
		"data":                  Web.Data,
		"Background":            Web.Style.Background,
		"AColor":                Web.Style.AColor,
		"AColorHover":           Web.Style.AColorHover,
	})
}

func NoRouter(ctx *gin.Context) {
	ctx.String(http.StatusOK, "The visit address does not exist, please return to the homepage")
}

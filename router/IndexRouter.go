// -*- coding: utf-8 -*-
// @Time    : 2021/11/15 17:10
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : IndexRouter.go

package router

import (
	"Gin_HomeNavigation/controllers"
	"github.com/gin-gonic/gin"
)

func IndexRouter(index *gin.Engine) {
	// 首页
	index.GET("/", controllers.Index)
}
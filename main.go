// -*- coding: utf-8 -*-
// @Time    : 2021/11/15 15:31
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : main.go

package main

import (
	"Gin_HomeNavigation/dataSource"
	"Gin_HomeNavigation/router"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 控制 Debug / Release 版本
	gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.DebugMode)

	// 创建
	r := gin.Default()

	//加载静态文件
	{
		// HTML文件
		r.LoadHTMLFiles("/app/views/index.html")
		// CSS
		r.Static("css", "/app/views/css")
		// JS
		r.Static("js", "/app/views/js")
		// IMG
		r.Static("img", "/app/views/img")

	}

	//Session中间件
	store := cookie.NewStore([]byte("userLogin"))
	r.Use(sessions.Sessions("pt_Ket", store))

	// 前端路由
	router.IndexRouter(r)

	// 读取配置文件
	conf := dataSource.LoadConfig()
	if conf.SoftWare.Port == "" {
		// 未配置端口,将使用默认端口: 8000
		log.Println("The port is not configured, the default port will be used: 8000")

		// 启动
		err := r.Run("0.0.0.0:8000")
		if err != nil {
			//启动时发生错误
			log.Println("An error occurred during startup：", err)
			panic(err)
		}
	} else {
		//即将开始监听并使用设定端口
		log.Println("About to start monitoring and use the set port：" + conf.SoftWare.Port)

		// 启动
		err := r.Run("0.0.0.0:" + conf.SoftWare.Port)
		if err != nil {
			//启动时发生错误
			log.Println("An error occurred during startup：", err)
			panic(err)
		}
	}
}

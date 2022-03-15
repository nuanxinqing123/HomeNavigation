// -*- coding: utf-8 -*-
// @Time    : 2022/3/14 17:55
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : routes.go

package routes

import (
	"Gin_HomeNavigation/bindata"
	"Gin_HomeNavigation/controllers"
	"Gin_HomeNavigation/logger"
	"Gin_HomeNavigation/middleware"
	"html/template"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	//r := gin.Default()
	r := gin.New()

	// 配置Gin日志
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 前端组
	{
		// 加载模板文件
		t, err := loadTemplate()
		if err != nil {
			panic(err)
		}
		r.SetHTMLTemplate(t)

		// 加载静态文件
		fs := assetfs.AssetFS{
			Asset:     bindata.Asset,
			AssetDir:  bindata.AssetDir,
			AssetInfo: nil,
			Prefix:    "assets",
		}
		r.StaticFS("/static", &fs)

		r.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", nil)
		})
	}

	// 公共权限组
	open := r.Group("api")
	{
		// 登录路径
		open.GET("f", controllers.LoginF)

		// 登录
		open.POST("login", controllers.Login)

		// 无密码数据
		open.GET("data", controllers.WebData)

		// Login数据
		open.GET("ldata", controllers.LoginData)
	}

	// 用户权限组
	api := r.Group("v1/api")
	api.Use(middleware.UserAuth())
	{
		api.GET("data", controllers.PwdWebData)
	}

	r.NoRoute(controllers.NoRouter)
	return r
}

//加载模板文件
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for _, name := range bindata.AssetNames() {
		if !strings.HasSuffix(name, ".html") {
			continue
		}
		asset, err := bindata.Asset(name)
		if err != nil {
			continue
		}
		name := strings.Replace(name, "assets/", "", 1)
		t, err = t.New(name).Parse(string(asset))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

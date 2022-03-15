// -*- coding: utf-8 -*-
// @Time    : 2021/11/15 15:31
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : main.go

package main

import (
	"Gin_HomeNavigation/dataSource"
	"Gin_HomeNavigation/logger"
	"Gin_HomeNavigation/routes"
	"Gin_HomeNavigation/tools/validator"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 加载配置文件
	conf := dataSource.LoadConfig()

	// 初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("Logger init failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("Logger success init ...")

	// 初始化翻译器
	if err := validator.InitTrans("zh"); err != nil {
		fmt.Printf("validator init failed, err:%v\n", err)
		return
	}

	// 运行模式
	if conf.SoftWare.Debug == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 注册路由
	r := routes.Setup()

	// 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.SoftWare.Port),
		Handler: r,
	}

	// 优雅关机
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listten: %s\n", err)
		}
	}()

	// 等待终端信号来优雅关闭服务器，为关闭服务器设置5秒超时
	quit := make(chan os.Signal, 1) // 创建一个接受信号的通道

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞此处，当接受到上述两种信号时，才继续往下执行
	zap.L().Info("准备关闭服务器")

	// 创建五秒超时的Context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 五秒内优雅关闭服务（将未处理完成的请求处理完再关闭服务），超过五秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("超时关闭服务器：", zap.Error(err))
	}

	zap.L().Info("服务器已关闭")
}

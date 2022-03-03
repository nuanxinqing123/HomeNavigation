// -*- coding: utf-8 -*-
// @Time    : 2021/11/15 17:10
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : Config_Conn.go

package dataSource

import (
	"Gin_HomeNavigation/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func LoadConfig() *model.Config {
	// 创建对象
	config := model.Config{}

	// 打开文件
	file ,err := os.Open("/app/conf/config.json")
	if err != nil {
		// 打开文件时发生错误
		log.Println("An error occurred while opening the file")
		panic(err)
	}
	// 延迟关闭
	defer file.Close()

	// 配置读取
	byteData, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		// 读取配置时发生错误
		log.Println("An error occurred while reading the configuration")
		panic(err)
	}

	// 数据绑定
	err3 := json.Unmarshal(byteData, &config)
	if err3 != nil {
		// 数据绑定时发生错误
		log.Println("An error occurred during data binding")
		panic(err)
	}

	return &config
}
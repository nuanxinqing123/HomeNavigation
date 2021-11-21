// -*- coding: utf-8 -*-
// @Time    : 2021/11/15 17:10
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : Config_Model.go

package model

// Config 配置项
type Config struct {
	Config   conf
	SoftWare software
	Data     []Data
	Style    style
}

type conf struct {
	Logo                  string `json:"logo"`
	Title                 string `json:"title"`
	Favicon               string `json:"favicon"`
	FooterTitle           string `json:"FooterTitle"`
	FooterTitleLink       string `json:"FooterTitleLink"`
	FooterTextDataOne     string `json:"FooterTextDataOne"`
	FooterTextDataTwo     string `json:"FooterTextDataTwo"`
	FooterTextDataTwoLink string `json:"FooterTextDataTwoLink"`
}

type software struct {
	Port string `json:"port"`
}

type Data struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Ico   string `json:"ico"`
	WLink string `json:"w_link"`
	NLink string `json:"n_link"`
}

type style struct {
	Background  string `json:"background"`
	AColor      string `json:"AColor"`
	AColorHover string `json:"AColorHover"`
}

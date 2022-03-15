// -*- coding: utf-8 -*-
// @Time    : 2022/3/14 18:08
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : code.go

package response

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota

	CodeInvalidParam = 2000 + iota

	CodeServerBusy
	CodeInvalidPageRequested

	CodeInvalidToken
	CodeNeedLogin
	CodePassWordWrong
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess: "Success",

	CodeInvalidParam:         "请求参数错误",
	CodeServerBusy:           "服务繁忙",
	CodeInvalidPageRequested: "请求无效页面",

	CodeInvalidToken:  "无效的Token",
	CodeNeedLogin:     "未登录",
	CodePassWordWrong: "密码错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

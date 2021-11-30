package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetSession 获取Session
func GetSession(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	loginuser := session.Get("userLogin")

	if loginuser != nil {
		return true
	} else {
		return false
	}
}

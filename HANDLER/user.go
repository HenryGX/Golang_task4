package HANDLER

import (
	"Golang_task4/SERVICE"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserCtrl struct{}

// 用户注册
func (u *UserCtrl) Register(c *gin.Context) {
	err := SERVICE.UserRegister(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// 用户登录
func (u *UserCtrl) Login(c *gin.Context) {
	err := SERVICE.UserLogin(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

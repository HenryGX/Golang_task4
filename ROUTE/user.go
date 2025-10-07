package ROUTE

import (
	"Golang_task4/HANDLER"

	"github.com/gin-gonic/gin"
)

// 用户路由
func UserRoutes(r *gin.Engine) {
	userCtrl := &HANDLER.UserCtrl{}
	u := r.Group("/user")
	{
		u.POST("/register", userCtrl.Register)
		u.POST("/login", userCtrl.Login)
	}
}

package ROUTE

import (
	"Golang_task4/HANDLER"
	"Golang_task4/MIDDLEWARE"

	"github.com/gin-gonic/gin"
)

// 评论路由
func CommentsRoutesInit(r *gin.Engine) {
	commentCtrl := &HANDLER.CommentsCtrl{}
	comment := r.Group("/comments", MIDDLEWARE.IsLoginMiddleware())
	{
		comment.POST("/add", commentCtrl.Create)
		comment.GET("/query/:id", commentCtrl.Query)
	}
}

package ROUTE

import (
	"Golang_task4/HANDLER"
	"Golang_task4/MIDDLEWARE"

	"github.com/gin-gonic/gin"
)

// 文章路由
func PostRoutes(r *gin.Engine) {
	postCtrl := &HANDLER.PostCtrl{}
	p := r.Group("/post", MIDDLEWARE.IsLoginMiddleware())
	{

		p.POST("/add", postCtrl.Add)
		p.GET("/query/:id", postCtrl.Query)
		p.GET("/queryList", postCtrl.QueryList)
		p.POST("/edit", postCtrl.Edit)
		p.POST("/delete", postCtrl.Delete)
	}
}

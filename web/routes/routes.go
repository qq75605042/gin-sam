package routes

import (
	"gin-sam/bootstrap"
	"gin-sam/web/controllers"
	"gin-sam/web/middleware"
	"gin-sam/web/routes/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Configure(b *bootstrap.Bootstrapper) {
	b.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello! Gin",
		})
	})

	// api分组
	apiGroup := b.Group("/api")

	apiGroup.GET("/auth", api.GetAuth)

	// middleware
	apiGroup.Use(middleware.JWT())
	{
		apiGroup.GET("/user/:id", controllers.GetUser)
	}

	// 加载模板路径
	//b.GET("/my_test", api.MyToken)
	b.LoadHTMLGlob("views/*")

	// web路由组
	webGroup := b.Group("/wap")

	webGroup.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Yookie",
		})
	})
}

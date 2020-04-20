package routes

import (
	"gin-sam/bootstrap"
	"gin-sam/web/routes/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Configure(b *bootstrap.Bootstrapper)  {
	b.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"message" : "Hello! Gin",
		})
	})

	b.GET("/auth",api.GetAuth)
	b.GET("/my_test",api.MyToken)
}

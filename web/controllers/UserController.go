package controllers

import (
	"gin-sam/utils/app"
	"gin-sam/utils/e"
	"net/http"

	"github.com/unknwon/com"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	user, err := userFService.Get(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user)
}

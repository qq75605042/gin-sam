package controllers

import (
	"net/http"

	"github.com/unknwon/com"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	user, e := userFService.Get(id)
	if e != nil {
		//todo...
	}
	c.JSON(http.StatusOK, user)
	return
}

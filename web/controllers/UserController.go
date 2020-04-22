package controllers

import (
	"gin-sam/services"
	"net/http"

	"github.com/unknwon/com"

	"github.com/gin-gonic/gin"
)

var UserService services.UserService

func init() {
	UserService = services.NewUserService()
}

func GetUser(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	user, e := UserService.Get(id)
	if e != nil {
		//todo...
	}
	c.JSON(http.StatusOK, user)
	return
}

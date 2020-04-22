package api

import (
	"fmt"
	"gin-sam/comm"
	"gin-sam/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//var validate *validator.Validate

//var config = &validator.Config{TagName: "validate"}

var validate = validator.New()

type auth struct {
	UserName string `form:"username" binding:"required" validate:"required"`
	Password string `form:"password" binding:"required" validate:"required"`
}

type JWTToken struct {
	Token    string `json:"token"`
	ExpireAt string `json:"expire_at"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	a := auth{
		UserName: username,
		Password: password,
	}

	errs := validate.Struct(a)
	if errs != nil {
		for _, v := range errs.(validator.ValidationErrors) {
			_ = fmt.Sprintf("[validate error] field:%s;tag:%s", v.Field(), v.Tag())
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    422,
				"message": fmt.Sprintf("field %s is %s", v.Field(), v.Tag()),
			})
			return
		}
	} else {
		if a.UserName != "admin" || a.Password != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized.",
			})
			return
		}
		if token, err := utils.GenerateToken(a.UserName, a.Password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "login success",
				"token":   token,
			})
		}
	}

}

func MyToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, JWTToken{
		Token:    "asdfafdasdf",
		ExpireAt: comm.FormatFromUnixTime(0),
	})
}

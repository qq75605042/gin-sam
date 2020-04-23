package api

import (
	"gin-sam/comm"
	"gin-sam/utils"
	"gin-sam/utils/app"
	"gin-sam/utils/e"
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
	appG := app.Gin{C: c}
	username := c.Query("username")
	password := c.Query("password")
	a := auth{
		UserName: username,
		Password: password,
	}

	errs := validate.Struct(a)
	if errs != nil {
		for _, v := range errs.(validator.ValidationErrors) {
			//c.JSON(http.StatusBadRequest, gin.H{
			//	"code":    422,
			//	"message": fmt.Sprintf("[validate error] field:%s;tag:%s", v.Field(), v.Tag()),
			//})
			appG.Response(http.StatusUnprocessableEntity, e.INVALID_PARAMS, v)
			return
		}
	} else {
		if a.UserName != "admin" || a.Password != "123456" {
			appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
			return
		}
		if token, err := utils.GenerateToken(a.UserName, a.Password); err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, err)
			return
		} else {
			appG.Response(http.StatusOK, e.SUCCESS, JWTToken{
				Token: token,
				//ExpireAt: "",
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

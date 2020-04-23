package middleware

import (
	"gin-sam/utils"
	"gin-sam/utils/app"
	"gin-sam/utils/e"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		appG := app.Gin{C: context}
		code := e.SUCCESS
		var data interface{}
		token := context.GetHeader("Authorization")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}
		if code != e.SUCCESS {
			appG.Response(http.StatusUnauthorized, code, data)
			context.Abort()
			return
		}
		context.Next()
	}
}

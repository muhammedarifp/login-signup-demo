package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammedarifp/signup/helpers"
)

func AuthWithToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr, terr := ctx.Cookie("token")

		if terr != nil {
			ctx.Redirect(302, "/login")
			ctx.Abort()
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return helpers.SecretKey, nil
		})

		if err != nil {
			ctx.Redirect(302, "/login")
			ctx.Abort()
			return
		}

		if !token.Valid {
			ctx.Redirect(302, "/login")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

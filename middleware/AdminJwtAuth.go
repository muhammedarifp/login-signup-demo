package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammedarifp/signup/helpers"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// tokenString := ctx.GetHeader("Authorization")
		tokenString, cerr := ctx.Cookie("ad-token")

		if cerr != nil {
			ctx.Redirect(302, "/admin/login")
			ctx.Abort()
			return
		}

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			// ctx.HTML(401, "unauth.html", gin.H{
			// 	"error": "Unauth user found",
			// })
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return helpers.SecretKey, nil
		})

		if err != nil || !token.Valid {
			ctx.HTML(401, "unauth.html", gin.H{
				"error": "Unauth user found",
			})
			return
		}

		ctx.Next()
	}
}

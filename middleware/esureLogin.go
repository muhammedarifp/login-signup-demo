package middleware

import (
	"github.com/gin-gonic/gin"
)

func EnsureUserLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := ctx.Cookie("token")
		if err == nil {
			ctx.Redirect(302, "/home")
		}
	}
}

func EnsureAdminLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := ctx.Cookie("ad-token")
		if err == nil {
			ctx.Redirect(302, "/admin/home")
		}
	}
}

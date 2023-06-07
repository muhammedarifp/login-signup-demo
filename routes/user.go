package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedarifp/signup/controllers"
	"github.com/muhammedarifp/signup/middleware"
)

func UserRoutes(g *gin.RouterGroup) {

	// User Get Points

	g.GET("/", middleware.AuthWithToken(), controllers.DefaultUroute())
	g.GET("/login", middleware.EnsureUserLogin(), controllers.UserLoginController())
	g.GET("/signup", middleware.EnsureAdminLogin(), controllers.UserSignUpController())
	g.GET("/home", middleware.AuthWithToken(), controllers.UserHomePage())
	g.GET("/logout-user", middleware.AuthWithToken(), controllers.UserLogoutController())

	// User Post Points

	g.POST("/signup", controllers.UserSignUpPostController())
	g.POST("/login", controllers.UserLoginPostController())
}

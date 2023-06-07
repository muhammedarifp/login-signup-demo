package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedarifp/signup/controllers"
	"github.com/muhammedarifp/signup/middleware"
)

func AdminRouter(g *gin.RouterGroup) {

	// Admnin Get Points

	g.GET("/", middleware.EnsureAdminLogin(), controllers.DefaultAdminController())
	g.GET("/login", middleware.EnsureAdminLogin(), controllers.LoginAdminRouteController)
	g.GET("/home", middleware.AdminAuthMiddleware(), controllers.AdminHomeController)
	g.GET("/remove-user/:id", middleware.AdminAuthMiddleware(), controllers.AdminRemoveUserController)
	g.GET("/add-user", middleware.AdminAuthMiddleware(), controllers.AddNewUser)
	g.GET("/update-user/:id", middleware.AdminAuthMiddleware(), controllers.AdminUpdateUserController)
	g.GET("/search-user", middleware.AdminAuthMiddleware(), controllers.SearchUserController)

	// Admin Post Points
	//

	g.POST("/add-user", controllers.AdminAddUserPostController)
	g.POST("/login", controllers.AdminLoginPostController)
	g.POST("/update-user", controllers.AdminUpdateUserPostController)
}

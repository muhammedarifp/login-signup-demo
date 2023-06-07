package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedarifp/signup/database"
	"github.com/muhammedarifp/signup/public"
	"github.com/muhammedarifp/signup/routes"
)

func main() {
	database.InitDataBase()
	database.DataBase.AutoMigrate(&public.Users{})
	router := gin.Default()
	router.LoadHTMLGlob("./static/html/*html")

	user := router.Group("/")
	routes.UserRoutes(user)

	admin := router.Group("/admin")
	routes.AdminRouter(admin)

	router.Run()
}

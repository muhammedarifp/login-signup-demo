package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedarifp/signup/database"
	"github.com/muhammedarifp/signup/helpers"
	"github.com/muhammedarifp/signup/public"
)

// This is a initial position
func DefaultAdminController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := *database.GetDb()
		var users []public.Users
		db.Find(&users)
		ctx.HTML(200, "admin-home.html", gin.H{
			"Users": users,
		})
	}
}

func LoginAdminRouteController(ctx *gin.Context) {
	ctx.HTML(200, "admin-login.html", gin.H{})
}

// Admin Get controller
func AdminHomeController(ctx *gin.Context) {
	db := *database.GetDb()
	var users []public.Users
	db.Find(&users)
	ctx.HTML(200, "admin-home.html", gin.H{
		"Users": users,
	})
}

func AddNewUser(ctx *gin.Context) {

	ctx.HTML(200, "admin-adduser.html", gin.H{})

}

func AdminAddUserPostController(ctx *gin.Context) {

	username := ctx.PostForm("name")
	email := ctx.PostForm("email")
	pass := ctx.PostForm("password")
	hashPass, _ := helpers.ToHash(pass)
	newUser := public.Users{
		Name:     username,
		Email:    email,
		Password: hashPass,
	}
	db := *database.GetDb()
	res := db.Create(&newUser)
	if res.Error == nil {
		ctx.Redirect(302, "/admin/home")
	}

}

func AdminLoginPostController(ctx *gin.Context) {
	email := ctx.PostForm("email")
	pass := ctx.PostForm("password")
	db := *database.GetDb()
	var admin public.Admins
	result := db.First(&admin, `email = ?`, email)
	if result.Error != nil {
		ctx.Redirect(302, "/admin/login")
	} else {
		val := helpers.CompareHash(admin.Password, pass)
		if !val {
			fmt.Println("not val error")
			ctx.Redirect(302, "/admin/login")
		} else {
			token, err := helpers.CreateJwtToken(admin.Name, admin.Email)
			if err != nil {
				log.Fatal("token not created")
			}
			ctx.SetCookie("ad-token", token, 3600, "/", "", false, true)
			ctx.Header("Authorization", token)
			ctx.Redirect(http.StatusFound, "/admin/home")
		}
	}

}

func AdminRemoveUserController(ctx *gin.Context) {

	userid := ctx.Param("id")
	db := *database.GetDb()
	// intVal, _ := strconv.Atoi(userid)
	res := db.Delete(&public.Users{}, userid)
	if res.Error != nil {
		fmt.Println("Erroe delete")
	} else {
		ctx.Redirect(302, "/admin/home")
	}

}

func AdminUpdateUserController(ctx *gin.Context) {

	id := ctx.Param("id")
	db := *database.GetDb()
	var user public.Users
	db.First(&user, id)
	ctx.HTML(200, "admin-updateuser.html", gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})

}

func AdminUpdateUserPostController(ctx *gin.Context) {

	id := ctx.PostForm("id")
	email := ctx.PostForm("email")
	name := ctx.PostForm("name")
	pass := ctx.PostForm("pass")
	db := *database.GetDb()
	var user public.Users
	if res := db.First(&user, id); res.Error != nil {
		log.Fatal("Db get value error")
		return
	}
	user.Email = email
	user.Name = name
	if pass != "" {
		passHash, err := helpers.ToHash(pass)
		if err != nil {
			log.Fatal("Pass word hash error")
			return
		}

		user.Password = passHash
	}

	db.Save(&user)
	ctx.Redirect(302, "/admin/home")

}

func SearchUserController(ctx *gin.Context) {

	val := ctx.Query("val")
	db := *database.GetDb()
	var users []public.Users
	db.Find(&users, `name=?`, val)
	ctx.HTML(200, "search-user.html", gin.H{
		"Users":  users,
		"Result": len(users) >= 1,
		"Count":  len(users),
	})

}

package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/muhammedarifp/signup/database"
	"github.com/muhammedarifp/signup/helpers"
	"github.com/muhammedarifp/signup/public"
)

// Fist / route. this is stating one
func DefaultUroute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(200, "home.html", gin.H{})
	}
}

// Signup get controller
func UserSignUpController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(200, "signup.html", gin.H{})
	}
}

// Login get controller
func UserLoginController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		ctx.Header("Pragma", "no-cache")
		ctx.HTML(200, "login.html", gin.H{})
	}
}

// Home page get controller
func UserHomePage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(200, "home.html", gin.H{})
	}
}

// User logout controller
func UserLogoutController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.SetCookie("token", "", -1, "/", "", false, true)
		ctx.Redirect(302, "/login")
	}
}

//................. // User Post methods // ................

type UserSignUpForm struct {
	Name  string `form:"name"`
	Email string `form:"email"`
	Pass1 string `form:"pass1"`
	Pass2 string `form:"pass2"`
}

// Signup post controller
func UserSignUpPostController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var formData UserSignUpForm
		ctx.ShouldBindWith(&formData, binding.Form)
		fmt.Println(formData)
		if len(formData.Name) >= 6 && formData.Email != "" && formData.Pass1 != "" && formData.Pass1 == formData.Pass2 {

			// Password Hashing ..!
			hashed, err := helpers.ToHash(formData.Pass1)
			if err != nil {
				log.Fatal(err.Error())
			}

			// Create New User
			NewData := public.Users{
				Name:     formData.Name,
				Email:    formData.Email,
				Password: hashed,
			}

			// Db instance
			db := *database.GetDb()

			// Insert data into database
			res := db.Create(&NewData)
			if res.Error != nil {
				log.Fatal(res.Error.Error())
			} else {
				ctx.Redirect(302, "/home")
			}
		} else {
			// if any eror redirect to signup with error messages
			ctx.HTML(200, "signup.html", gin.H{
				"error": "Something error try again",
			})
		}
	}
}

// Login Post controller
func UserLoginPostController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.PostForm("email")
		pass := ctx.PostForm("password")
		db := *database.GetDb()
		var result public.Users
		res := db.First(&result, `email = ?`, email)
		if res.Error != nil {
			ctx.HTML(200, "login.html", gin.H{
				"error": "email incorrect..! check your email",
			})
			return
		}
		passStatus := helpers.CompareHash(result.Password, pass)
		fmt.Println(passStatus)
		if passStatus {
			token, _ := helpers.CreateJwtToken(result.Name, result.Email)
			ctx.SetCookie("token", token, 86400, "/", "localhost:8080", false, true)
			ctx.Redirect(302, "/home")
		} else {
			ctx.HTML(200, "login.html", gin.H{
				"error": "password incorrect..! check your password",
			})
		}

	}
}

package main

import (
	"github.com/ajujacob88/go-jwt-auth-basic-webapp/controllers"
	"github.com/ajujacob88/go-jwt-auth-basic-webapp/initializers"
	"github.com/ajujacob88/go-jwt-auth-basic-webapp/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	//fmt.Println("Hello3")

	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()

	/*
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.Run() // listen and serve on 0.0.0.0:3000 ,, http://localhost:3000/ping  ,, since 3000 given in env
	*/
}

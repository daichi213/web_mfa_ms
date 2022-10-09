package api

import (
	"fmt"
	"log"
    "github.com/gin-gonic/gin"
	"api/controllers"
	"api/middleware"
)

func InitializeRoutes(router *gin.Engine) {
	fmt.Println("checkpoint 1")
	// Call the authMiddleware
	authMiddleware, err := middleware.CallAuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// Sign Up
	router.POST("/signup", controllers.SignUp)

	// Login
	router.POST("/login", authMiddleware.LoginHandler)

	// 404のRouting
	router.NoRoute(authMiddleware.MiddlewareFunc(), controllers.NoRouting)

	// 認証後のRouting
	auth := router.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", controllers.HelloHandler)
		// USER関連
		auth.PUT("/:user_id", controllers.UserEditHandler)
		auth.DELETE("/:user_id", controllers.UserDeleteHandler)
	}

	// AuthMiddleWareの初期化
	if errInit := authMiddleware.MiddlewareInit();errInit != nil {
		log.Fatal("AuthMiddleware.MiddlewareInit failed: ", errInit.Error())
	}
}
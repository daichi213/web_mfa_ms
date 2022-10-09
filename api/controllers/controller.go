package controllers

import (
	"log"
    "github.com/gin-gonic/gin"

	jwt "github.com/appleboy/gin-jwt/v2"
	"api/middleware"
	"api/models"
)

func NoRouting(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(404, gin.H{"code":"PAGE_NOT_FOUND", "message": "Page not found"})
}

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(middleware.IdentityKey)
	c.JSON(200, gin.H{
		"userID": claims[middleware.IdentityKey],
		"userName": user.(*models.Login).UserName,
		"text": "Hello World.",
	})
}

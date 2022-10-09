package controllers

import (
	"log"
	"net/http"
    "github.com/gin-gonic/gin"
	"api/models"

	// jwt "github.com/appleboy/gin-jwt/v2"
)

func SignUp(c *gin.Context) {
	var signupUser models.Login
	err := c.BindJSON(&signupUser)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Printf("BindJSON is failed: %v", err)
	} else {
		err := models.CreateUser(&signupUser)
		if err != nil {
			// TODO エラー発生時にルートパスへリダイレクトさせる処理を追加する
			c.Status(http.StatusBadRequest)
			log.Printf("CreateUser is failed: %v", err)
		} else {
			c.JSON(200, gin.H{
				"UserName": signupUser.UserName,
				"Email": signupUser.Email,
			})
			// c.Redirect(http.StatusFound, "/auth/schedule")
		}
	}
}

func UserEditHandler(c *gin.Context) {
	// TODO ユーザーの登録情報編集
}

func UserDeleteHandler(c *gin.Context) {
	// TODO ユーザーの登録情報編集
}
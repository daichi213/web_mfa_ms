package middleware

import (
	"fmt"
	"time"
	"log"
    "github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"api/models"

	jwt "github.com/appleboy/gin-jwt/v2"
)

// TODO MODELの変数をmiddlewareでも参照してしまっているため、分離できるように努める（DIコンテナで実現できるか？）
// jwt middleware
var IdentityKey = "email"

func CallAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	fmt.Println("checkpoint 2")
	AuthMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:		"test zone",
		Key:  		[]byte("secret key"),
		Timeout:	time.Hour,
		MaxRefresh:	time.Hour,
		IdentityKey: IdentityKey,
		// login後に呼び出される関数
		PayloadFunc: InternalPayloadFunc,
		// Authorizatorへ値を渡すための関数
		IdentityHandler: InternalIdentityHandlerFunction,
		// 認証(ユーザー本人かどうかの確認)
		Authenticator: InternalAuthenticatorFunction,
		// 認可(権限の確認)
		// token発行後のページの読み込み制御についての関数
		Authorizator: InternalAuthorizatorFunction,
		Unauthorized: InternalUnauthorizedFunction,
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})
	return AuthMiddleware, err
}

func InternalPayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.Login); ok {
		return jwt.MapClaims{
			IdentityKey: v.Email,
		}
	}
	return jwt.MapClaims{}
}

// 認可後に実行される関数
func InternalIdentityHandlerFunction(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	// return &Login{
	// 	Email: claims[IdentityKey].(string),
	// }
	if err := models.GetUserByEmail(claims[IdentityKey].(string)); err != nil {
		log.Fatalf("GetUserByEmail is failed: %v", err)
		return nil
	} else {
		return &models.UserFromDB
	}
}

// 認証関数
func InternalAuthenticatorFunction(c *gin.Context) (interface{}, error) {
	var loginVals models.Login
	if err := c.ShouldBind(&loginVals); err != nil {
		fmt.Println("checkpoint shouldbind")
		return "", jwt.ErrMissingLoginValues
	}

	// TODOHTTPヘッダからIPアドレスを記録できるようにする
	if err := models.GetUserByEmail(loginVals.Email); err != nil {
		// log.Fatalf("No existing password is sent")
		fmt.Println("checkpoint get user")
		return "", jwt.ErrMissingLoginValues
	}

	if invalid := bcrypt.CompareHashAndPassword(models.UserFromDB.Password, []byte(loginVals.Password)); invalid != nil {
		log.Fatalf("Password is wrong...;dbside:%v,loginVals:%v", models.UserFromDB.Password, []byte(loginVals.Password))
		return "", jwt.ErrFailedAuthentication
	} else {
		return &models.Login{
			UserName: 	models.UserToDB.UserName,
			Email: 		models.UserToDB.Email,
			Password:	loginVals.Password,
		}, nil
	}
}

// 認可関数
func InternalAuthorizatorFunction(data interface{}, c *gin.Context) bool {
	if _, ok := data.(*models.Login); ok {
		return true
	}
	return false
}

// 認可失敗後に実行される関数
func InternalUnauthorizedFunction(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":		code,
		"message":	message,
	})
}
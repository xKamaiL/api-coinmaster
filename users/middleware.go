package users

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/xkamail/api-ts3-gin/common"
	"net/http"
)

func stripBearerPrefixFromTokenString(token string) (string, error) {
	if len(token) > 7 && token[0:7] == "Bearer " {
		return token[7:], nil
	}
	return token, nil
}

var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromTokenString,
}

var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func UpdateContextUserModel(c *gin.Context, myUserId uint) {
	if myUserId != 0 {

		repository := NewUserRepository()
		user, err := repository.FindById(myUserId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "User Not Found.",
			})
		} else {
			fmt.Println("[Initialization] User Data")
			c.Set("user_id", myUserId)
			c.Set("user", user)
			c.Next()
		}
	}

}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// set for zero
		UpdateContextUserModel(c, 0)
		token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := []byte(common.JWTSecretConfig)
			return b, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "No authenticate",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			myUserId := uint(claims["sub"].(float64))
			UpdateContextUserModel(c, myUserId)
		}
	}
}

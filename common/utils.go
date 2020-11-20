package common

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"reflect"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// A helper function to generate random string
func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

const JWTSecretConfig = "EGaXs0pqbNQbfPnR3I58sgE6J0XCDzDN09q9kGfg9wtamErPp9Mb6JWIXEooR8MO"

// A Util function to generate jwt_token which can be used in the request header
func GenerateJWTToken(id int) string {
	jwtToken := jwt.New(jwt.GetSigningMethod("HS256"))
	jwtToken.Claims = jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	}

	token, _ := jwtToken.SignedString([]byte(JWTSecretConfig))

	return token
}

// My own Error type that will help return my customized Error info
//  {"database": {"hello":"no such table", error: "not_exists"}}
type Error struct {
	Errors  map[string]interface{} `json:"errors"`
	Message string                 `json:"message"`
}

// To handle the error returned by c.Bind in gin framework
func NewError(key string, err error) Error {
	res := Error{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

func NewErrorMessage(message string) Error {
	res := Error{}
	res.Message = message
	return res
}

func NewValidatorError(err error) Error {
	response := Error{}
	response.Errors = make(map[string]interface{})
	if reflect.TypeOf(err).String() == "*errors.errorString" {
		response.Errors["errors"] = err.Error()
		response.Message = "Error"
		return response
	}
	errs := err.(validator.ValidationErrors)
	response.Message = "Validate Failed!"
	for _, v := range errs {
		if v.Param() != "" {
			response.Errors[v.Field()] = fmt.Sprintf("%v %v", v.Tag(), v.Param())
		} else {
			response.Errors[v.Field()] = fmt.Sprintf("key %v", v.Tag())
		}

	}

	return response
}

func AppUrl() string {
	return "https://api-new.ts.in.th/"
}

// Changed the c.MustBindWith() ->  c.ShouldBindWith().
// I don't want to auto return 400 when error happened.
// origin function is here: https://github.com/gin-gonic/gin/blob/master/context.go
func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

func Message(message string) gin.H {
	return gin.H{
		"message": message,
	}
}

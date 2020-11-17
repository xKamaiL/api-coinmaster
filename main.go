package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xkamail/api-coinmaster/users"
	"net/http"
	"strings"
)

type appError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		if len(detectedErrors) > 0 {
			fmt.Println("some error")
			err := detectedErrors[0].Err
			var parsedError *appError
			switch err.(type) {
			default:
				parsedError = &appError{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}
			c.AbortWithStatusJSON(parsedError.Code, parsedError)
			return
		}

	}
}

func JSONAppErrorReporter() gin.HandlerFunc {
	return jsonAppErrorReporterT(gin.ErrorTypeAny)
}
func setupRouter() *gin.Engine {
	r := gin.Default()
	// using error report into json.
	r.Use(JSONAppErrorReporter())
	r.Use(cors.New(cors.Config{
		AllowMethods:    []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:    []string{"*"},
		AllowFiles:      true,
		AllowAllOrigins: true,
	}))

	// make a public path link to /uploads
	//r.Static("/uploads", "./public")

	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	v1 := r.Group("/api")
	authorized := r.Group("/api")
	authorized.Use(users.AuthMiddleware())
	// setup router app.
	users.GuestRouter(v1.Group("/user"))
	users.AuthenticateRouter(authorized.Group("/user"))

	return r
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
}

func main() {

	initConfig()

	if viper.GetString("app.env") == "production" {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("[Production] Enabled")
	}

	port := viper.GetString("app.port")
	if port == "" {
		port = "8080"
	}
	db := common.Init()
	defer db.Close()

	r := setupRouter()
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	})
	_ = r.Run(":" + port)
}

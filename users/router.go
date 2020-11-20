package users

import "github.com/gin-gonic/gin"

func GuestRouter(r *gin.RouterGroup) {
	r.POST("/login", UserLogin)
}

func AuthenticateRouter(r *gin.RouterGroup) {
	r.GET("/me")
}

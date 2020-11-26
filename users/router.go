package users

import "github.com/gin-gonic/gin"

func GuestRouter(r *gin.RouterGroup) {
	r.POST("/login", UserLogin)
	r.POST("/register", UserRegister)
}

func AuthenticateRouter(r *gin.RouterGroup) {
	r.GET("/me", UserProfile)
}

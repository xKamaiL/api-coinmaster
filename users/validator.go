package users

import (
	"github.com/gin-gonic/gin"
	"github.com/xkamail/api-coinmaster/common"
)

type LoginValidator struct {
	Username string `form:"username" json:"username" binding:"required,alphanum,min=4,max=255"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=255"`
}

func NewLoginValidator() LoginValidator {
	return LoginValidator{}
}

func (v *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}

	return nil
}

type RegisterValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"required,alphanum,min=4,max=255"`
		Password string `form:"password" json:"password" binding:"required,min=6,max=255"`
		Email    string `form:"email" binding:"required,max=255,email"`
		MobileNo string `form:"email" binding:"required,number,len=10"`
	}
	userModel User `json:"-"`
}

func NewRegisterValidator() RegisterValidator {
	return RegisterValidator{}
}

func (v *RegisterValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	v.userModel.Username = v.User.Username
	v.userModel.Password, _ = common.HashPassword(v.User.Password)
	v.userModel.IpAddress = c.ClientIP()
	v.userModel.Email = v.User.Email
	v.userModel.Role = "customer"
	v.userModel.SpinPoint = 0
	v.userModel.InvitePoint = 0

	return nil
}

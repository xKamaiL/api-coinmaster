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

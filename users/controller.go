package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xkamail/api-coinmaster/common"
	"net/http"
	"time"
)

func UserLogin(c *gin.Context) {
	repository := NewUserRepository()
	body := NewLoginValidator()
	err := body.Bind(c)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	user, err := repository.FindByUsername(body.Username)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, common.Message("ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง"))
		return
	}
	check := common.CheckPasswordHash(body.Password, user.Password)
	if !check {
		c.JSON(400, common.Message("ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง"))
		return
	}
	user.IpAddress = c.ClientIP()
	user.UpdatedAt = time.Now()
	go repository.Save(user)

	c.JSON(http.StatusOK, gin.H{
		"message": "เข้าสู่ระบบสำเร็จ",
		"token":   common.GenerateJWTToken(user.ID),
	})
	return

}

func UserRegister(c *gin.Context) {

}

func UserProfile(c *gin.Context) {

}

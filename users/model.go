package users

import "time"

type User struct {
	ID              int       `gorm:"column:id",json:"id"`
	Username        string    `gorm:"column:username",json:"username"`
	Role            string    `gorm:"column:role",json:"role"`
	SpinPoint       int       `gorm:"column:spin_point",json:"spin_point"`
	Email           string    `gorm:"column:email",json:"email"`
	MobileNo        string    `gorm:"column:mobile_no",json:"mobile_no"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at",json:"email_verified_at"`
	Password        string    `gorm:"column:password",json:"password"`
	RememberToken   string    `gorm:"column:remember_token",json:"remember_token"`
	IpAddress       string    `gorm:"column:ip_address",json:"ip_address"`
	CreatedAt       time.Time `gorm:"column:created_at",json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at",json:"updated_at"`
}

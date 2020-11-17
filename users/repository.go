package users

import "github.com/xkamail/api-coinmaster/common"

type Repository interface {
	FindAll() ([]User, error)
	FindById(id uint) (User, error)
	First(condition interface{}) (User, error)
}

func NewUserRepository() Repository {
	return &UserRepository{}
}

type UserRepository struct{}

func (u UserRepository) FindAll() ([]User, error) {
	var users []User
	err := common.GetDB().Find(&users).Error
	return users, err
}

func (u UserRepository) FindById(id uint) (User, error) {
	panic("implement me")
}

func (u UserRepository) First(condition interface{}) (User, error) {
	panic("implement me")
}

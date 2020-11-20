package users

import "github.com/xkamail/api-coinmaster/common"

type Repository interface {
	FindAll() ([]User, error)
	FindById(id uint) (User, error)
	First(condition interface{}) (User, error)
	FindByUsername(username string) (User, error)
	FindByEmail(email string) (User, error)
	Save(user User) (User, error)
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
	var user User
	db := common.GetDB()
	err := db.Where(condition).First(&user).Error
	return user, err
}

func (u UserRepository) FindByUsername(username string) (User, error) {
	var user User
	db := common.GetDB()
	err := db.Where(User{Username: username}).First(&user).Error
	return user, err
}

func (u UserRepository) FindByEmail(email string) (User, error) {
	var user User
	db := common.GetDB()
	err := db.Where(User{Email: email}).First(&user).Error
	return user, err
}

func (u UserRepository) Save(user User) (User, error) {
	err := common.GetDB().Save(&user).Error
	return user, err
}
//
//func (u UserRepository) Delete(user User)  {
//
//}

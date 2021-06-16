package v1

import (
	"github.com/jinzhu/gorm"
	"vipulsirdemo/models"
)

type UserRepository interface {
	UserSignUp(user models.User)
	UserLogin(user models.User)
	UserVerification(user models.User)
}

type UserRepo struct {
	DB *gorm.DB
}

func (ur *UserRepo) UserSignUp(user models.User) {
	ur.DB.Model(&user).Create(&user)
	return
}

func (ur *UserRepo) GetUserByEmail(email string) (models.User, error) {
	user := models.User{}
	err := ur.DB.Model(user).Select("name, email, image, gender, mobile, hobby").Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepo) GetUserById(id uint) (models.User, error) {
	user := models.User{}
	err := ur.DB.Model(user).Select("name, email, image, gender, mobile, hobby").Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepo) UserLogin(user models.User) (models.User, error) {
	err := ur.DB.Model(user).Select("id, name, email, image, gender, mobile, hobby").Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepo) EditProfile(user models.User) {
	ur.DB.Model(&user).Update(&user)
	return
}

func (ur *UserRepo) UserVerification(user models.User) {
	ur.DB.Model(&user).Create(&user)
	return
}

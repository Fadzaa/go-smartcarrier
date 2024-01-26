package auth

import "gorm.io/gorm"

//Initialize interface

type Repository interface {
	RegisterUser(user User) (User, error)
	LoginUser(user User) (User, error)
}

// Initialize struct

type RepositoryImpl struct {
	db *gorm.DB
}

// Implement struct value

func NewUserRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

//Method from userRepository struct that implement Repository api

func (r *RepositoryImpl) RegisterUser(user User) (User, error) {

	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *RepositoryImpl) LoginUser(user User) (User, error) {

	err := r.db.Where("email = ?", user.Email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

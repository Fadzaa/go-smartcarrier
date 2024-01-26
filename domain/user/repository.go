package user

import "gorm.io/gorm"

//Initialize api

type UserRepository interface {
	FindAll() ([]User, error)
	FindUserByID(id int) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id int) (User, error)
}

// Initialize struct
type UserRepositoryImpl struct {
	db *gorm.DB
}

// Implement struct value

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

//Method from userRepository struct that implement UserRepository api

func (r *UserRepositoryImpl) FindAll() ([]User, error) {
	var users []User
	//SELECT * FROM users
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) FindUserByID(id int) (User, error) {
	var user User
	//SELECT * FROM users WHERE id = ?
	err := r.db.First(&user, id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) CreateUser(user User) (User, error) {

	//INSERT INTO users (name, email, occupation) VALUES (?, ?, ?)
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user User) (User, error) {
	//UPDATE users SET name=?, email=?, occupation=? WHERE id=?
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) DeleteUser(id int) (User, error) {
	var user User
	//DELETE FROM users WHERE id=?
	err := r.db.Delete(&user, id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

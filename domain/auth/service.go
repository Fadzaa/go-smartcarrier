package auth

// Init Service

type Service interface {
	RegisterUser(user User) (User, error)
	LoginUser(user User) (User, error)
}

// Init struct

type ServiceImpl struct {
	repositoryUser Repository
}

// Implement struct value

func NewAuthService(repositoryUser Repository) *ServiceImpl {
	return &ServiceImpl{repositoryUser}
}

// Method from ServiceImpl struct that implement Service api

func (s *ServiceImpl) RegisterUser(user User) (User, error) {
	newUser, err := s.repositoryUser.RegisterUser(User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *ServiceImpl) LoginUser(user User) (User, error) {
	newUser, err := s.repositoryUser.LoginUser(User{
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

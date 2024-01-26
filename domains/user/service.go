package user

// Init Service
type UserService interface {
	GetAllUser() ([]User, error)
	GetUserByID(id int) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id int) (User, error)
}

// Init struct
type UserServiceImpl struct {
	repositoryUser UserRepository
}

// Implement struct value
func NewUserService(repositoryUser UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repositoryUser}
}

// Method from UserServiceImpl struct that implement UserService interface
func (s *UserServiceImpl) GetAllUser() ([]User, error) {
	users, err := s.repositoryUser.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *UserServiceImpl) GetUserByID(id int) (User, error) {
	user, err := s.repositoryUser.FindUserByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *UserServiceImpl) CreateUser(user User) (User, error) {
	newUser, err := s.repositoryUser.CreateUser(User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *UserServiceImpl) UpdateUser(user User) (User, error) {
	user, err := s.repositoryUser.UpdateUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *UserServiceImpl) DeleteUser(id int) (User, error) {
	user, err := s.repositoryUser.DeleteUser(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

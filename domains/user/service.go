package user

// Init Service
type ServiceUser interface {
	GetAllUser() ([]User, error)
	GetUserByID(id int) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id int) (User, error)
}

// Init struct
type serviceUser struct {
	repositoryUser RepositoryUser
}

// Implement struct value
func NewUserService(repositoryUser RepositoryUser) *serviceUser {
	return &serviceUser{repositoryUser}
}

// Method from serviceUser struct that implement ServiceUser interface
func (s *serviceUser) GetAllUser() ([]User, error) {
	users, err := s.repositoryUser.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *serviceUser) GetUserByID(id int) (User, error) {
	user, err := s.repositoryUser.FindUserByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *serviceUser) CreateUser(user User) (User, error) {
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

func (s *serviceUser) UpdateUser(user User) (User, error) {
	user, err := s.repositoryUser.UpdateUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *serviceUser) DeleteUser(id int) (User, error) {
	user, err := s.repositoryUser.DeleteUser(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

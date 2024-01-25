package user

type ServiceUser interface {
	GetAllUser() ([]User, error)
	GetUserByID(id int) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id int) (User, error)
}

type serviceUser struct {
	repositoryUser RepositoryUser
}

func NewUserService(repositoryUser RepositoryUser) *serviceUser {
	return &serviceUser{repositoryUser}
}

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
	user, err := s.repositoryUser.CreateUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
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

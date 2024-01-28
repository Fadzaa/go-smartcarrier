package course

type Service interface {
	GetAllCourse() ([]Course, error)
	GetCourseByID(id int) (Course, error)
	CreateCourse(job Course) (Course, error)
	UpdateCourse(job Course) (Course, error)
	DeleteCourse(id int) (Course, error)
}

type ServiceImpl struct {
	repositoryCourse Repository
}

func NewCourseService(repositoryCourse Repository) *ServiceImpl {
	return &ServiceImpl{repositoryCourse}
}

func (s *ServiceImpl) GetAllCourse() ([]Course, error) {
	courses, err := s.repositoryCourse.FindAll()
	if err != nil {
		return courses, err
	}

	return courses, nil
}

func (s *ServiceImpl) GetCourseByID(id int) (Course, error) {
	course, err := s.repositoryCourse.FindCourseByID(id)
	if err != nil {
		return course, err
	}

	return course, nil
}

func (s *ServiceImpl) CreateCourse(job Course) (Course, error) {
	newCourse, err := s.repositoryCourse.CreateCourse(Course{
		Title: job.Title,
	})
	if err != nil {
		return newCourse, err
	}

	return newCourse, nil
}

func (s *ServiceImpl) UpdateCourse(course Course) (Course, error) {
	course, err := s.repositoryCourse.UpdateCourse(course)
	if err != nil {
		return course, err
	}

	return course, nil
}

func (s *ServiceImpl) DeleteCourse(id int) (Course, error) {
	course, err := s.repositoryCourse.DeleteCourse(id)
	if err != nil {
		return course, err
	}

	return course, nil
}

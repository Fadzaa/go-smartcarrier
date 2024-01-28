package course

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Course, error)
	FindCourseByID(id int) (Course, error)
	CreateCourse(job Course) (Course, error)
	UpdateCourse(job Course) (Course, error)
	DeleteCourse(id int) (Course, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

func (r *RepositoryImpl) FindAll() ([]Course, error) {
	var courses []Course
	err := r.db.Find(&courses).Error
	if err != nil {
		return courses, err
	}

	return courses, nil
}

func (r *RepositoryImpl) FindCourseByID(id int) (Course, error) {
	var course Course
	err := r.db.First(&course, id).Error
	if err != nil {
		return course, err
	}

	return course, nil
}

func (r *RepositoryImpl) CreateCourse(course Course) (Course, error) {

	err := r.db.Create(&course).Error
	if err != nil {
		return course, err
	}

	return course, nil
}

func (r *RepositoryImpl) UpdateCourse(course Course) (Course, error) {

	err := r.db.Save(&course).Error
	if err != nil {
		return course, err
	}

	return course, nil
}

func (r *RepositoryImpl) DeleteCourse(id int) (Course, error) {
	var course Course
	err := r.db.Delete(&course, id).Error
	if err != nil {
		return course, err
	}

	return course, nil
}

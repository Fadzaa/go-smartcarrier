package job

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Job, error)
	FindJobByID(id int) (Job, error)
	CreateJob(job Job) (Job, error)
	UpdateJob(job Job) (Job, error)
	DeleteJob(id int) (Job, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

func (r *RepositoryImpl) FindAll() ([]Job, error) {
	var jobs []Job
	err := r.db.Find(&jobs).Error
	if err != nil {
		return jobs, err
	}

	return jobs, nil
}

func (r *RepositoryImpl) FindJobByID(id int) (Job, error) {
	var job Job
	err := r.db.First(&job, id).Error
	if err != nil {
		return job, err
	}

	return job, nil
}

func (r *RepositoryImpl) CreateJob(job Job) (Job, error) {

	err := r.db.Create(&job).Error
	if err != nil {
		return job, err
	}

	return job, nil
}

func (r *RepositoryImpl) UpdateJob(job Job) (Job, error) {

	err := r.db.Save(&job).Error
	if err != nil {
		return job, err
	}

	return job, nil
}

func (r *RepositoryImpl) DeleteJob(id int) (Job, error) {
	var job Job
	err := r.db.Delete(&job, id).Error
	if err != nil {
		return job, err
	}

	return job, nil
}

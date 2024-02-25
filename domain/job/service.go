package job

type Service interface {
	GetAllJob() ([]Job, error)
	GetJobByID(id int) (Job, error)
	CreateJob(job Job) (Job, error)
	UpdateJob(job Job) (Job, error)
	DeleteJob(id int) (Job, error)
}

type ServiceImpl struct {
	repositoryJob Repository
}

func NewJobService(repositoryJob Repository) *ServiceImpl {
	return &ServiceImpl{repositoryJob}
}

func (s *ServiceImpl) GetAllJob() ([]Job, error) {
	jobs, err := s.repositoryJob.FindAll()
	if err != nil {
		return jobs, err
	}

	return jobs, nil
}

func (s *ServiceImpl) GetJobByID(id int) (Job, error) {
	job, err := s.repositoryJob.FindJobByID(id)
	if err != nil {
		return job, err
	}

	return job, nil
}

func (s *ServiceImpl) CreateJob(job Job) (Job, error) {
	newJob, err := s.repositoryJob.CreateJob(job)
	if err != nil {
		return newJob, err
	}

	return newJob, nil
}

func (s *ServiceImpl) UpdateJob(job Job) (Job, error) {
	job, err := s.repositoryJob.UpdateJob(job)
	if err != nil {
		return job, err
	}

	return job, nil
}

func (s *ServiceImpl) DeleteJob(id int) (Job, error) {
	job, err := s.repositoryJob.DeleteJob(id)
	if err != nil {
		return job, err
	}

	return job, nil
}

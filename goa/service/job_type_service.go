package service

import (
	"goa/goa/model"
	"goa/goa/repo"
)

type JobTypeService interface {
	GetAll() []model.JobType
	GetByID(id int64) (model.JobType, bool)
	DeleteByID(id int64) bool
	Save(jobType model.JobType) (model.JobType, error)
}

type jobTypeService struct {
	repo repo.JobTypeRepo
}

func NewJobTypeService(r repo.JobTypeRepo) JobTypeService {
	return &jobTypeService{
		repo: r,
	}
}

func (s *jobTypeService) GetAll() []model.JobType {
	return s.repo.FindAll()
}
func (s *jobTypeService) GetByID(id int64) (model.JobType, bool) {

	return s.repo.FindOne(id)
}
func (s *jobTypeService) DeleteByID(id int64) bool {
	return s.repo.DeleteOne(id)
}
func (s *jobTypeService) Save(jobType model.JobType) (model.JobType, error) {

	return s.repo.Save(jobType)
}

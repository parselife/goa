package service

import (
	"goa/goa/model"
	"goa/goa/repo"
)

type ProService interface {
	GetAll() []model.JobProject
	GetByID(id int64) (model.JobProject, bool)
	DeleteByID(id int64) bool
	Save(pro model.JobProject) (model.JobProject, error)
}

type proService struct {
	repo repo.ProjectRepository
}

func NewProService(repo repo.ProjectRepository) ProService {
	return &proService{repo: repo}
}

func (s *proService) GetAll() []model.JobProject {
	return s.repo.FindAll()
}
func (s *proService) GetByID(id int64) (model.JobProject, bool) {
	return s.repo.FindOne(id)
}
func (s *proService) DeleteByID(id int64) bool {
	return s.repo.DeleteOne(id)
}
func (s *proService) Save(pro model.JobProject) (model.JobProject, error) {
	return s.repo.Save(pro)
}

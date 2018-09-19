package service

import (
	"goa/goa/model"
	"goa/goa/repo"
)

type OrganService interface {
	GetAll() []model.Organ
	GetByID(id int64) (model.Organ, bool)
	DeleteByID(id int64) bool
	Save(organ model.Organ) (model.Organ, error)
}

type organService struct {
	repo repo.OrganRepo
}

func NewOrganService(r repo.OrganRepo) OrganService {
	return &organService{
		repo: r,
	}
}

func (s *organService) GetAll() []model.Organ {
	return s.repo.FindAll()
}

func (s *organService) GetByID(id int64) (model.Organ, bool) {
	return s.repo.FindOne(id)
}

func (s *organService) DeleteByID(id int64) bool {
	return s.repo.DeleteOne(id)
}

func (s *organService) Save(organ model.Organ) (model.Organ, error) {
	return s.repo.Save(organ)
}

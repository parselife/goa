package service

import (
	"goa/goa/model"
	"goa/goa/repo"
)

type JobLogService interface {
	GetAll() []model.JobLog
	GetByID(id int64) (model.JobLog, bool)
	GetByUserId(userId int64) ([]model.JobLog, bool)

	DeleteByID(id int64) bool

	Save(jobLog model.JobLog) (model.JobLog, error)
}

type jobLogService struct {
	repo repo.JobLogRepository
}

func NewJobLogService(r repo.JobLogRepository) JobLogService {
	return &jobLogService{r}
}

func (s *jobLogService) GetAll() []model.JobLog {
	return s.repo.FindAll()
}

func (s *jobLogService) GetByID(id int64) (model.JobLog, bool) {

	return s.repo.FindOne(id)

}

func (s *jobLogService) GetByUserId(userId int64) ([]model.JobLog, bool) {

	return s.repo.FindByUserId(userId)
}

func (s *jobLogService) DeleteByID(id int64) bool {
	return s.repo.DeleteOne(id)
}

func (s *jobLogService) Save(jobLog model.JobLog) (model.JobLog, error) {

	return s.repo.Save(jobLog)

}

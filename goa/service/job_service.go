package service

import (
	"goa/goa/model"
	"goa/goa/repo"
	"time"
)

type JobLogService interface {
	GetAll() []model.JobLog
	GetByID(id int64) (model.JobLog, bool)
	GetByUserId(userId int64) ([]model.JobLog, bool)

	DeleteByID(id int64) bool
	Save(jobLog model.JobLog) (model.JobLog, error)

	GetWeekly() []model.JobLog
	GetMonthly() []model.JobLog
	GetDaily() []model.JobLog
	GetThreeDays() []model.JobLog
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

// 获取当前一周的工作日志
func (s *jobLogService) GetWeekly() []model.JobLog {
	now := time.Now()
	now.Day()

	return nil

}

// 获取当月的工作日志
func (s *jobLogService) GetMonthly() []model.JobLog {

	return nil
}

// 获取当天的工作日志
func (s *jobLogService) GetDaily() []model.JobLog {
	return nil
}

// 获取昨天今天明天的工作日志
func (s *jobLogService) GetThreeDays() []model.JobLog {
	return nil
}

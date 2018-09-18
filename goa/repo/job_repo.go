package repo

import "goa/goa/model"

type JobLogRepository interface {
	FindOne(id int64) (jobLog model.JobLog, found bool)
	FindByUserId(userId int64) ([]model.JobLog, bool)
	FindByOrganId(organId int64) ([]model.JobLog, bool)
	FindAll() ([]model.JobLog)

	Save(jobLog model.JobLog) (updated model.JobLog, err error)
	DeleteOne(id int64) (ok bool)
}

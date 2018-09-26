package repo

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"goa/goa/model"
)

type JobLogRepository interface {
	FindOne(id int64) (jobLog model.JobLog, found bool)
	FindByUserId(userId int64) ([]model.JobLog, bool)
	FindByOrganId(organId int64) ([]model.JobLog, bool)
	FindAll() []model.JobLog

	Save(jobLog model.JobLog) (updated model.JobLog, err error)
	DeleteOne(id int64) (ok bool)
}

func NewJobLogRepo(engine *xorm.Engine) JobLogRepository {
	return &jobLogRepo{orm: engine}
}

type jobLogRepo struct {
	orm *xorm.Engine
}

func (j *jobLogRepo) FindOne(id int64) (jobLog model.JobLog, found bool) {
	ret := new(model.JobLog)
	has, err := j.orm.Id(id).Get(ret)
	if err != nil {
		fmt.Println(err)
		return *ret, false
	}
	return *ret, has
}

func (j *jobLogRepo) FindAll() []model.JobLog {
	joblogs := make([]model.JobLog, 0)
	err := j.orm.Asc("start_time").Find(&joblogs)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return joblogs
}

func (j *jobLogRepo) FindByUserId(userId int64) ([]model.JobLog, bool) {
	jobs := make([]model.JobLog, 0)
	err := j.orm.Where("user = ?", userId).Asc("start_time").Find(&jobs)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return jobs, true
}

func (j *jobLogRepo) FindByOrganId(organId int64) ([]model.JobLog, bool) {

	jobs := make([]model.JobLog, 0)
	err := j.orm.Where("organ_id = ?", organId).Asc("start_time").Find(&jobs)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return jobs, true
}

func (j *jobLogRepo) Save(jobLog model.JobLog) (updated model.JobLog, err error) {
	if jobLog.ID > 0 {
		// update
		_, err = j.orm.Where("i_d = ?", jobLog.ID).Update(&jobLog)
		if err != nil {
			return jobLog, err
		}
		return jobLog, nil
	} else {
		//	insert
		_, err = j.orm.Insert(&jobLog)
		if err != nil {
			fmt.Println(err)
			return jobLog, err
		}
		return jobLog, nil
	}

}

func (j *jobLogRepo) DeleteOne(id int64) (ok bool) {
	_, err := j.orm.Delete(&model.JobLog{ID: id})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

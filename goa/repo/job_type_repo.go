package repo

import (
	"goa/goa/model"
	"github.com/go-xorm/xorm"
	"fmt"
)

type JobTypeRepo interface {
	FindOne(id int64) (model.JobType, bool)
	FindAll() []model.JobType

	Save(jobType model.JobType) (model.JobType, error)
	DeleteOne(id int64) bool
}

func NewJobTypeRepo(engine *xorm.Engine) JobTypeRepo {
	return &jobTypeRepo{
		orm: engine,
	}
}

type jobTypeRepo struct {
	orm *xorm.Engine
}

func (r *jobTypeRepo) FindOne(id int64) (model.JobType, bool) {
	ret := new(model.JobType)
	has, err := r.orm.Id(id).Get(ret)
	if err != nil {
		fmt.Println(err)
		return *ret, false
	}
	return *ret, has
}

func (r *jobTypeRepo) FindAll() []model.JobType {
	rets := make([]model.JobType, 0)
	err := r.orm.Find(&rets)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return rets

}

func (r *jobTypeRepo) Save(jobType model.JobType) (model.JobType, error) {
	if jobType.ID > 0 {
		// update
		_, err := r.orm.Where("i_d = ?", jobType.ID).Update(&jobType)
		if err != nil {
			fmt.Println(err)
			return jobType, err
		}
		return jobType, nil
	} else {
		//	insert
		_, err := r.orm.Insert(&jobType)
		if err != nil {
			fmt.Println(err)
			return jobType, err
		}
		return jobType, nil
	}
}

func (r *jobTypeRepo) DeleteOne(id int64) bool {
	_, err := r.orm.Delete(&model.JobType{ID: id})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

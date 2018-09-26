package repo

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"goa/goa/model"
)

type OrganRepo interface {
	FindOne(id int64) (model.Organ, bool)
	FindAll() []model.Organ

	Save(organ model.Organ) (model.Organ, error)
	DeleteOne(id int64) bool
}

func NewOrganRepo(engine *xorm.Engine) OrganRepo {
	return &organRepo{
		orm: engine,
	}
}

type organRepo struct {
	orm *xorm.Engine
}

func (r *organRepo) FindOne(id int64) (model.Organ, bool) {
	ret := new(model.Organ)
	has, err := r.orm.Id(id).Get(ret)
	if err != nil {
		fmt.Println(err)
		return *ret, false
	}
	return *ret, has

}
func (r *organRepo) FindAll() []model.Organ {
	rets := make([]model.Organ, 0)
	err := r.orm.Find(&rets)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return rets
}

func (r *organRepo) Save(organ model.Organ) (model.Organ, error) {
	if organ.ID > 0 {
		// update
		_, err := r.orm.Where("i_d = ?", organ.ID).Update(&organ)
		if err != nil {
			fmt.Println(err)
			return organ, err
		}
		return organ, nil
	} else {
		//	insert
		_, err := r.orm.Insert(&organ)
		if err != nil {
			fmt.Println(err)
			return organ, err
		}
		return organ, nil
	}
}
func (r *organRepo) DeleteOne(id int64) bool {
	_, err := r.orm.Delete(&model.Organ{ID: id})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

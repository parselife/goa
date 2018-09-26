package repo

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"goa/goa/model"
)

type ProjectRepository interface {
	FindOne(id int64) (pro model.JobProject, found bool)
	FindAll() []model.JobProject

	Save(pro model.JobProject) (model.JobProject, error)
	DeleteOne(id int64) (ok bool)
	FindByStatus(status model.ProjectStatus) []model.JobProject
}

type projectRepo struct {
	orm *xorm.Engine
}

func NewProjectRepo(db *xorm.Engine) ProjectRepository {
	return &projectRepo{orm: db}
}

func (p *projectRepo) FindOne(id int64) (pro model.JobProject, found bool) {
	ret := new(model.JobProject)
	has, err := p.orm.Id(id).Get(ret)
	if err != nil {
		fmt.Println(err)
		return *ret, false
	}
	return *ret, has
}

func (p *projectRepo) FindAll() (users []model.JobProject) {
	rets := make([]model.JobProject, 0)
	err := p.orm.Find(&rets)
	if err != nil {
		fmt.Println(err)
		return
	}
	return rets
}

func (p *projectRepo) Save(pro model.JobProject) (updatedUser model.JobProject, err error) {
	if pro.ID > 0 {
		// update
		_, err = p.orm.Where("i_d = ?", pro.ID).Update(&pro)
		if err != nil {
			fmt.Println(err)
			return pro, err
		}
		return pro, nil
	} else {
		//	insert
		_, err := p.orm.Insert(&pro)
		if err != nil {
			fmt.Println(err)
			return pro, err
		}
		return pro, nil
	}
}

func (p *projectRepo) DeleteOne(id int64) (ok bool) {
	_, err := p.orm.Delete(&model.JobProject{ID: id})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (p *projectRepo) FindByStatus(status model.ProjectStatus) []model.JobProject {
	pros := make([]model.JobProject, 0)
	err := p.orm.Where("status = ?", status).Find(&pros)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pros
}

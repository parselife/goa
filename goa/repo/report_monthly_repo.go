package repo

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"goa/goa/model"
)

type MonthlyReportRepository interface {
	FindOne(id int64) (m model.MonthlyReport, found bool)
	FindAll() []model.MonthlyReport

	Save(r model.MonthlyReport) (model.MonthlyReport, error)
	DeleteOne(id int64) (ok bool)
}

func NewMonthlyReportRepo(engine *xorm.Engine) MonthlyReportRepository {
	return &monthlyReportRepo{orm: engine}
}

type monthlyReportRepo struct {
	orm *xorm.Engine
}

func (m *monthlyReportRepo) FindOne(id int64) (model.MonthlyReport, bool) {
	ret := new(model.MonthlyReport)
	return *ret, false
}

func (m *monthlyReportRepo) FindAll() []model.MonthlyReport {
	rets := make([]model.MonthlyReport, 0)
	err := m.orm.Find(&rets)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return rets
}

func (m *monthlyReportRepo) Save(r model.MonthlyReport) (model.MonthlyReport, error) {
	if r.ID > 0 {
		// update
		_, err := m.orm.Where("i_d = ?", r.ID).Update(&r)
		if err != nil {
			fmt.Println(err)
			return r, err
		}
		return r, nil
	} else {
		//	insert
		_, err := m.orm.Insert(&r)
		if err != nil {
			fmt.Println(err)
			return r, err
		}
		return r, nil
	}
}

func (m *monthlyReportRepo) DeleteOne(id int64) bool {
	_, err := m.orm.Delete(&model.MonthlyReport{ID: id})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

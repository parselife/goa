package repo

import "github.com/go-xorm/xorm"

type MonthlyReportRepository interface {
	// todo
}

func NewMonthlyReportRepo(engine *xorm.Engine) MonthlyReportRepository {
	return &monthlyReportRepo{orm: engine}
}

type monthlyReportRepo struct {
	orm *xorm.Engine
}

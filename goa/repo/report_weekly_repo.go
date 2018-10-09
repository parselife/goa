package repo

import "github.com/go-xorm/xorm"

type WeeklyReportRepository interface {
	// todo
}

func NewWeeklyReportRepo(engine *xorm.Engine) WeeklyReportRepository {
	return &weeklyReportRepo{orm: engine}
}

type weeklyReportRepo struct {
	orm *xorm.Engine
}

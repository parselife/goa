package model

import (
	"goa/goa/core"
	"time"
)

// 月报
type MonthlyReport struct {
	ID         int64     `json:"id"`
	User       User      `json:"user" xorm:"not null"`                 // 关联用户
	ReportTime time.Time `json:"reportTime" xorm:"not null"`           // 汇报时间
	Summary    string    `json:"summary" xorm:"varchar(500) not null"` // 工作总结
	Undone     string    `json:"undone" xorm:"varchar(400)"`           // 未完成事项
	Consult    string    `json:"consult" xorm:"varchar(400)"`          // 需要领导协调事项
	CreateAt   core.Time `json:"createAt" xorm:"created"`
	UpdateAt   core.Time `json:"updateAt" xorm:"updated"`
}

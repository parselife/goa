package model

import (
	"goa/goa/core"
	"time"
)

// 工时记录
type JobLog struct {
	ID        int64      `json:"id"`
	User      User       `json:"user" xorm:"not null"`
	Project   JobProject `json:"project" xorm:"not null"`
	Type      JobType    `json:"type" xorm:"not null"`
	Title     string     `json:"title" xorm:"varchar(50) not null"`
	Content   string     `json:"content" xorm:"varchar(1000) not null "`
	StartTime time.Time  `json:"startTime" xorm:"not null"`
	EndTime   time.Time  `json:"endTime" xorm:"not null"`
	CreateAt  core.Time  `json:"createAt" xorm:"created"`
	UpdateAt  core.Time  `json:"updateAt" xorm:"updated"`
}

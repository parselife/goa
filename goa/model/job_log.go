package model

import (
	"time"
)

// 工时记录
type JobLog struct {
	ID        int64      `json:"id"`
	UserId    int64      `json:"userId"`
	Project   JobProject `json:"project" xorm:"not null"`
	Type      JobType    `json:"type" xorm:"not null"`
	Title     string     `json:"title" xorm:"varchar(50) not null"`
	Content   string     `json:"content" xorm:"varchar(1000) not null "`
	StartTime time.Time  `json:"startTime"`
	EndTime   time.Time  `json:"endTime"`
	CreateAt  time.Time  `json:"createAt" xorm:"created"`
	UpdateAt  time.Time  `json:"updateAt" xorm:"updated"`
}

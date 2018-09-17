package model

import "time"

// 工时记录
type JobLog struct {
	ID        int64      `json:"id"`
	UserId    int64      `json:"user_id"`
	Project   JobProject `json:"project"`
	Type      JobType    `json:"type"`
	Title     string     `json:"title"`
	Content   string     `json:"content, omitempty" xorm:"varchar(1000)"`
	StartTime time.Time  `json:"start_time"`
	EndTime   time.Time  `json:"end_time"`
	CreateAt  time.Time  `json:"create_at" xorm:"created"`
	UpdateAt  time.Time  `json:"update_at" xorm:"updated"`
}

package model

import "time"

type ProjectStatus int32

const (
	running   ProjectStatus = iota
	paused
	completed
)

// 工时关联的项目
type JobProject struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Status      ProjectStatus `json:"status"`
	Description string        `json:"desc,omitempty" xorm:"varchar(200)"`
	CreateAt    time.Time     `json:"create_at" xorm:"created"`
	UpdateAt    time.Time     `json:"update_at" xorm:"updated"`
}

package model

import (
	"goa/goa/core"
)

type ProjectStatus int32

const (
	running   ProjectStatus = iota
	paused
	completed
)

// 工时关联的项目
type JobProject struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name" xorm:"varchar(50) notnull unique"`
	Status      ProjectStatus `json:"status"`
	Description string        `json:"desc,omitempty"`
	CreateAt    core.Time     `json:"createAt" xorm:"created"`
	UpdateAt    core.Time     `json:"updateAt" xorm:"updated"`
}

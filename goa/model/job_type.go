package model

import (
	"goa/goa/core"
)

// 工时类型
type JobType struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" xorm:"notnull unique"`
	Alias       string    `json:"alias" xorm:"notnull"`
	Description string    `json:"desc, omitempty" xorm:"varchar(200)"`
	CreateAt    core.Time `json:"createAt" xorm:"created"`
	UpdateAt    core.Time `json:"updateAt" xorm:"updated"`
}

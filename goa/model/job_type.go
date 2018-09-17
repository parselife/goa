package model

import "time"

// 工时类型
type JobType struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Alias       string    `json:"alias"`
	Description string    `json:"desc, omitempty" xorm:"varchar(200)"`
	CreateAt    time.Time `json:"create_at" xorm:"created"`
	UpdateAt    time.Time `json:"update_at" xorm:"updated"`
}

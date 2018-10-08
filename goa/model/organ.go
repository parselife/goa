package model

import "goa/goa/core"

// 员工部门
type Organ struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" xorm:"varchar(50) notnull unique"`
	Description string    `json:"desc,omitempty"`
	Manager     int64      `json:"manager" xorm:"notnull"` //部门主管id
	CreateAt    core.Time `json:"createAt" xorm:"created"`
	UpdateAt    core.Time `json:"updateAt" xorm:"updated"`
}

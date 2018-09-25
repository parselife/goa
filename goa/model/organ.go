package model

// 员工部门
type Organ struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" xorm:"varchar(50) notnull unique"`
	Description string `json:"desc,omitempty"`
}

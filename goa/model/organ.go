package model

// 部门
type Organ struct {
	ID          int64
	Name        string `json:"name" xorm:"notnull, unique"`
	Description string `json:"desc,omitempty" xorm:"varchar(255)"`
}

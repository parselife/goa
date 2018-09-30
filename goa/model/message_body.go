package model

import "goa/goa/core"

// 消息内容类型 可以是 普通文本/markdown 文本/html
type MessageType int32

const (
	text   MessageType = iota
	marked
)

// 消息体
type MessageBody struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title" xorm:"varchar(50) not null"`
	Type     MessageType `json:"type" xorm:"not null default 0"`
	Detail   string    `json:"detail" xorm:"varchar(255) not null"`
	CreateAt core.Time `json:"createAt" xorm:"created"`
}

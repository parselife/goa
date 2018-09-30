package model

import "goa/goa/core"

type MessageStatus int32

const (
	unread    MessageStatus = iota // 未读
	readed                         // 已读
	withdrawn                      // 被撤回 *撤回后的消息不会出现在收信人的信箱里 但会在 发信人的信箱里
)

// 应用消息
type Message struct {
	ID       int64         `json:"id"`
	Sender   User          `json:"sender"`   // 发信人
	Receiver User          `json:"receiver"` // 收信人
	Status   MessageStatus `json:"status" xorm:"not null (default 0)"`   //当前状态
	MessageBody MessageBody `json:"messageBody" xorm:"not null"`
	CreateAt core.Time     `json:"createAt" xorm:"created"`
	UpdateAt core.Time     `json:"updateAt" xorm:"updated"`
}

package repo

import "github.com/go-xorm/xorm"

type MsgRepository interface {

}

func NewMsgRepo(engine *xorm.Engine) MsgRepository {
	return &msgRepo{orm: engine}
}

type msgRepo struct {
	orm *xorm.Engine
}

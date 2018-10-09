package repo

import "github.com/go-xorm/xorm"

type MsgBodyRepository interface {

}

func NewMsgBodyRepo(engine *xorm.Engine) MsgBodyRepository {
	return &msgBodyRepo{orm: engine}
}

type msgBodyRepo struct {
	orm *xorm.Engine
}
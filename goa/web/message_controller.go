package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"goa/goa/core"
	"goa/goa/service"
)

type MessageController struct {
	Ctx     iris.Context
	Service service.MsgService
	Session *sessions.Session
}

//---------------------工作汇报-------------------
// GET /rest/msg/all
func (c *MessageController) GetAll() interface{} {

	return core.RenderFailure("xxx")
}

//


package web

import (
	"github.com/kataras/iris"
	"goa/goa/service"
	"github.com/kataras/iris/sessions"
)

// 管理员controller
// - 管理用户 查看所有用户工作记录
type AdminController struct {
	Ctx iris.Context
	UserService service.UserService
	Session *sessions.Session
}
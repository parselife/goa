package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"goa/goa/service"
)

type JobController struct {
	Ctx           iris.Context
	UserService   service.UserService
	JobLogService service.JobLogService
	Session       *sessions.Session
}

//func (c *RestController) authCheck() {
//	if !core.IsLoggedIn(c.Session) {
//		c.Ctx.Redirect("/login?url=" + c.Ctx.Request().Host + c.Ctx.Request().RequestURI)
//	}
//}

// GET /rest/jobs

package web

import (
	"github.com/kataras/iris"
	"goa/goa/service"
	"github.com/kataras/iris/sessions"
)

type ReportController struct {
	Ctx     iris.Context
	Service service.ReportService
	Session *sessions.Session
}

// GET /rest/report/monthly
func (c *RestController) GetMonthly() interface{} {

	//core.GetCurrentUserID(c.Session)
	return nil
}

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

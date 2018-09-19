package web


import (
	"github.com/kataras/iris"
	"goa/goa/service"
	"github.com/kataras/iris/sessions"
	"goa/goa/model"
	"goa/goa/core"
	"github.com/kataras/iris/mvc"
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
func (c *JobController) GetJobs() interface{} {
	c.authCheck()
	if !core.Admin(c.Session) {
		return mvc.Response{
			Code: iris.StatusForbidden,
		}
	}
	return c.JobLogService.GetAll()
}

// GET /rest/job/1
func (c *JobController) GetJobBy(id int64) interface{} {
	c.authCheck()
	o, found := c.JobLogService.GetByID(id)
	if !found {
		return core.RenderFailure("未找到记录!")
	}
	return core.RenderJson(o)
}

// POST /rest/job
func (c *JobController) PostJob() interface{} {
	c.authCheck()
	var joblog model.JobLog
	if err := c.Ctx.ReadJSON(&joblog); err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
	}
	saved, err := c.JobLogService.Save(joblog)
	if err != nil {
		return core.RenderFailure(err.Error())
	}
	return core.RenderJson(saved)
}

// DELETE /rest/job/1
func (c *JobController) DeleteJobBy(id int64) interface{} {
	c.authCheck()
	ok := c.OrganService.DeleteByID(id)
	return core.RenderJson(ok)
}
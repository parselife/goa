package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"goa/goa/core"
	"goa/goa/model"
	"goa/goa/service"
)

// rest controller
// - 管理用户 查看所有用户工作记录
type RestController struct {
	Ctx           iris.Context
	UserService   service.UserService
	JobLogService service.JobLogService
	ProService    service.ProService
	OrganService  service.OrganService
	TypeService   service.JobTypeService
	Session       *sessions.Session
}

func (c *RestController) authCheck() {
	if !core.IsLoggedIn(c.Session) {
		c.Ctx.Redirect("/login?url=" + c.Ctx.Request().Host + c.Ctx.Request().RequestURI)
	}
}

//---------------------用户-----------------------
// 获取所有用户 GET /rest/users
func (c *RestController) GetUsers() mvc.Response {
	c.authCheck()
	if !core.Admin(c.Session) {
		return mvc.Response{
			Code: iris.StatusForbidden,
		}
	}
	return core.RenderJson(c.UserService.GetAll())
}

// GET /rest/user/1
func (c *RestController) GetUserBy(id int64) interface{} {
	c.authCheck()
	user, ok := c.UserService.GetByID(id)
	if !ok {
		return core.RenderFailure("用户不存在")
	}
	return core.RenderJson(user)
}

// 删除用户 DELETE /rest/user/12
func (c *RestController) DeleteUserBy(id int64) mvc.Response {
	c.authCheck()
	if !core.Admin(c.Session) {
		return mvc.Response{
			Code: iris.StatusForbidden,
		}
	}
	ok := c.UserService.DeleteByID(id)
	return core.RenderJson(ok)
}

// 提交用户 POST /rest/user
func (c *RestController) PostUser() interface{} {
	c.authCheck()
	if !core.Admin(c.Session) {
		return mvc.Response{
			Code: iris.StatusForbidden,
		}
	}
	var u model.User
	if err := c.Ctx.ReadJSON(&u); err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
	}
	user, found := c.UserService.GetByID(u.ID)
	if found && u.Password != "" && user.Password != u.Password {
		u.Password = model.Md5Password(u.Password)
	}
	saved, err := c.UserService.Save(u)
	if err != nil {
		return core.RenderFailure(err.Error())
	}
	return core.RenderJson(saved)
}

//-----------------项目---------------------------
// 获取所有项目
func (c *RestController) GetPros() mvc.Response {
	c.authCheck()
	return core.RenderJson(c.ProService.GetAll())
}

// GET /rest/pro/12
func (c *RestController) GetProBy(id int64) interface{} {
	c.authCheck()
	p, found := c.ProService.GetByID(id)
	if !found {
		return core.RenderFailure("未找到记录!")
	}
	return core.RenderJson(p)
}

// 提交 POST /rest/pro
func (c *RestController) PostPro() interface{} {
	c.authCheck()
	if !core.Admin(c.Session) {
		return mvc.Response{
			Code: iris.StatusForbidden,
		}
	}
	var p model.JobProject
	if err := c.Ctx.ReadJSON(&p); err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
	}
	saved, err := c.ProService.Save(p)
	if err != nil {
		return core.RenderFailure(err.Error())
	}
	return core.RenderJson(saved)
}

// 删除项目 DELETE /rest/pro/12
func (c *RestController) DeleteProBy(id int64) mvc.Response {
	c.authCheck()
	if !core.Admin(c.Session) {
		return mvc.Response{
			Code: iris.StatusForbidden,
		}
	}
	ok := c.ProService.DeleteByID(id)
	return core.RenderJson(ok)
}

//--------------工作类型----------------
// GET /rest/types
func (c *RestController) GetTypes() interface{} {
	c.authCheck()
	return c.TypeService.GetAll()
}

// GET /rest/type/1
func (c *RestController) GetTypeBy(id int64) interface{} {
	c.authCheck()
	o, found := c.TypeService.GetByID(id)
	if !found {
		return core.RenderFailure("未找到记录!")
	}
	return core.RenderJson(o)
}

// POST /rest/type
func (c *RestController) PostType() interface{} {
	c.authCheck()
	if !core.Admin(c.Session) {
		return mvc.Response{
			Code: iris.StatusForbidden,
		}
	}
	var jobType model.JobType
	if err := c.Ctx.ReadJSON(&jobType); err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
	}
	saved, err := c.TypeService.Save(jobType)
	if err != nil {
		return core.RenderFailure(err.Error())
	}
	return core.RenderJson(saved)
}

// DELETE /rest/type/1
func (c *RestController) DeleteTypeBy(id int64) interface{} {
	c.authCheck()
	if !core.Admin(c.Session) {
		return mvc.Response{
			Code: iris.StatusForbidden,
		}
	}
	ok := c.TypeService.DeleteByID(id)
	return core.RenderJson(ok)
}

//--------------部门-------------------
// GET /rest/organs
func (c *RestController) GetOrgans() interface{} {
	c.authCheck()
	return c.OrganService.GetAll()
}

// GET /rest/organ/1
func (c *RestController) GetOrganBy(id int64) interface{} {
	c.authCheck()
	o, found := c.OrganService.GetByID(id)
	if !found {
		return core.RenderFailure("未找到记录!")
	}
	return core.RenderJson(o)
}

// POST /rest/organ
func (c *RestController) PostOrgan() interface{} {
	c.authCheck()
	if !core.Admin(c.Session) {
		return mvc.Response{
			Code: iris.StatusForbidden,
		}
	}
	var organ model.Organ
	if err := c.Ctx.ReadJSON(&organ); err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
	}
	saved, err := c.OrganService.Save(organ)
	if err != nil {
		return core.RenderFailure(err.Error())
	}
	return core.RenderJson(saved)
}

// DELETE /rest/organ/1
func (c *RestController) DeleteOrganBy(id int64) interface{} {
	c.authCheck()
	if !core.Admin(c.Session) {
		return mvc.Response{
			Code: iris.StatusForbidden,
		}
	}
	ok := c.OrganService.DeleteByID(id)
	return core.RenderJson(ok)
}

//func (c *RestController) BeforeActivation(b mvc.BeforeActivation) {
//	fmt.Println("before")
//
//}
//
//func (c *RestController) AfterActivation(a mvc.AfterActivation) {
//	if a.Singleton() {
//		panic("basicController should be stateless, a request-scoped, we have a 'Session' which depends on the context.")
//	}
//}

//func (c *RestController) All() {
//
//	fmt.Println("---all---")
//}
//
//func (c *RestController) Any() {
//
//	fmt.Println("---any----")
//
//}
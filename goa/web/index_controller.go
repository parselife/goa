package web

import (
	"github.com/kataras/iris"
	"goa/goa/service"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/mvc"
	"goa/goa/model"
	"goa/goa/core"
)

// 默认 controller
type IndexController struct {
	Ctx iris.Context

	UserService service.UserService

	// 由 main 注入
	Session *sessions.Session
}

func (c *IndexController) GetLogout() {
	c.Session.Set(core.AUTHENTICATED, false)
	c.Session.Destroy()
}

func (c *IndexController) GetText() string {
	return "hello alex"
}
// 首页
func (c *IndexController) Get() mvc.View {
	return mvc.View{
		Name: "/index.html",
		Data: iris.Map{"Title": "首页", "Author": "alex"},
	}
}


func (c *IndexController) GetLogin()mvc.Response {

	c.Session.Set(core.AUTHENTICATED, true)
	return mvc.Response {
		Code:200,
	}

}

// 登录
func (c *IndexController) PostLogin() mvc.Result {
	var (
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)
	u, ok := c.UserService.GetByUsername(username)
	if !ok {
		return mvc.Response{

		}
	}
	// 验证密码
	valid := model.Md5Password(password) == u.Password
	if !valid {
		return mvc.Response{
			Text: "",
		}
	}

	c.Session.Set(core.UserId, u.ID)
	c.Session.Set(core.AUTHENTICATED, true)
	c.Session.Set(core.IsAdmin, u.IsAdmin)

	return mvc.Response{
		Path: "/",
	}
}


//func (c *IndexController) BeforeActivation(b mvc.BeforeActivation) {
//	fmt.Println("before")
//
//}
//
//func (c *IndexController) AfterActivation(a mvc.AfterActivation) {
//	if a.Singleton() {
//		panic("basicController should be stateless, a request-scoped, we have a 'Session' which depends on the context.")
//	}
//}

//func (c *IndexController) All() {
//
//	fmt.Println("---all---")
//}
//
//func (c *IndexController) Any() {
//
//	fmt.Println("---any----")
//
//}

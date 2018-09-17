package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"goa/goa/service"
	"github.com/kataras/iris/mvc"
	"goa/goa/model"
)

type UserController struct {
	// context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding)
	// and the Session which depends on the current context (dynamic binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	Service service.UserService

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session
}

const (
	userIDKey = "UserID"
	authenticated = "Authenticated"
	isAdmin = "IsAdmin"
)


func (c *UserController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(userIDKey, 0)
	return userID
}

func (c *UserController) isLoggedIn() bool {
	return c.getCurrentUserID() > 0
}

// testcode
func (c *UserController) GetText() mvc.Response {
	return mvc.Response{
		ContentType: "application/json",
		Text: "{'name':'alex'}",
	}

}

func (c *UserController) logout() {
	c.Session.Destroy()
}

func (c *UserController) PostLogin() mvc.Response {
	var (
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)

	u, ok := c.Service.GetByUsername(username)

	if !ok {

		return mvc.Response{

		}
	}
	// validate password
	valid := model.Md5Password(password) == u.Password
	if !valid {
		return mvc.Response{
			Text: "",
		}
	}

	c.Session.Set(userIDKey, u.ID)
	c.Session.Set(authenticated, true)
	c.Session.Set(isAdmin, u.IsAdmin)

	return mvc.Response{
		Path: "/",
	}


}
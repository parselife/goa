package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"goa/goa/core"
	"goa/goa/service"
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

func (c *UserController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(core.UserId, 0)
	return userID
}

func (c *UserController) isLoggedIn() bool {
	return c.getCurrentUserID() > 0
}

// testcode
func (c *UserController) GetText() mvc.Response {
	if ok, _ := c.Session.GetBoolean(core.AUTHENTICATED); !ok {
		return mvc.Response{
			Code: iris.StatusUnauthorized,
			Text: "login first",
		}
	}
	return mvc.Response{
		ContentType: "application/json",
		Text:        "{'name':'alex'}",
	}

}

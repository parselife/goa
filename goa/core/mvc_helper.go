package core

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
)

func RenderFailure(msg string) mvc.Response {

	ret := iris.Map{
		"succcess": false,
		"msg":      msg,
	}

	return mvc.Response{
		Object: ret,
	}
}

func RenderJson(obj interface{}) mvc.Response {

	return mvc.Response{
		ContentType: "application/json",
		Object:      obj,
	}
}



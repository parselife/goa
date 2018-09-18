package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"goa/goa/core"
	"goa/goa/datasource"
	"goa/goa/repo"
	"goa/goa/service"
	"goa/goa/web"
	"log"
	"time"
	"fmt"
)

func main() {

	app := newApp()
	// 解析配置
	c := iris.TOML("./etc/app.tml")
	appConf := parseConfig(c.Other)
	// 初始化数据库连接
	db, err := datasource.Init(&appConf.SqlConf)
	if err != nil {
		app.Logger().Fatalf("db init error: %v", err)
		return
	}
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	// ---- 初始化 session ----
	sessionManager := sessions.New(sessions.Config{
		Cookie:       "njzykjgoa",
		Expires:      12 * time.Hour,
		AllowReclaim: true,
	})

	//app.Use(securityMiddleware)

	// ---- 依赖注入 ----
	userService := service.NewUserService(repo.NewUserRepo(db))
	user := mvc.New(app.Party("/user"))
	user.Register(
		userService,
		sessionManager.Start,
	).Handle(new(web.UserController))

	// 首页
	app.Get("/", func(ctx iris.Context) {
		// 检测是否登录
		if auth, _ := sessionManager.Start(ctx).GetBoolean("authenticated"); !auth {
			ctx.Redirect("/login")
			return
		}
		ctx.ViewData("Title", "首页")
		ctx.View("index.html")
	})
	//登录页
	app.Get("/login", func(ctx iris.Context) {
		ctx.View("login.html")
	})
	// 应用描述信息
	app.Get("/info", func(ctx iris.Context) {
		ctx.JSON(appConf.AppInfo)
	})

	app.Run(iris.Addr(":8080"), iris.WithConfiguration(c))
}

// todo
func securityMiddleware(ctx iris.Context) {

	fmt.Println(ctx.Path())
	fmt.Println(ctx.RequestPath(true))
	ctx.Next()
}

// 初始化app
func newApp() *iris.Application {
	app := iris.New()
	// 设置日志
	app.Logger().Prefix = []byte("[goa]")
	app.Logger().SetLevel("info")
	// 注册视图
	app.RegisterView(iris.HTML("./public", ".html").Reload(true))
	assetHandler := app.StaticHandler("./public", false, false)
	app.SPA(assetHandler)
	app.Logger().Info("app is ready")
	return app
}

// 解析应用配置
func parseConfig(conf map[string]interface{}) core.AppConf {
	sqlConfMap, ok := conf["sqlConf"].(map[string]interface{})
	if ok != true {
		log.Fatal("sql config error")
	}
	infoConfMap, ok := conf["appInfo"].(map[string]interface{})
	if ok != true {
		log.Fatal(" app config error")
	}
	return core.AppConf{core.SqlConf{sqlConfMap["DriverName"].(string), sqlConfMap["DbUrl"].(string),
		sqlConfMap["ShowSql"].(bool), sqlConfMap["LogLevel"].(int64)}, core.AppInfo{infoConfMap["Name"].(string),
		infoConfMap["Desc"].(string), infoConfMap["Author"].(string)}}

}

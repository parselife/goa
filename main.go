package main

import (
	"github.com/kataras/iris"
	"log"
	"goa/goa/core"
	"goa/goa/datasource"
	"goa/goa/repo"
	"goa/goa/service"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("info")
	app.RegisterView(iris.HTML("./public", ".html"))

	assetHandler := app.StaticHandler("./public", false, false)
	app.SPA(assetHandler)

	c := iris.TOML("./etc/app.tml")

	appConf := parseConfig(c.Other)

	db, err := datasource.Init(&appConf.SqlConf)
	if err !=nil {
		log.Fatal("db init error")
	}

	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)



	app.Get("/index", func(ctx iris.Context) {
		//ctx.ViewData("App", appConfig)
		ctx.View("index.html")
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(appConf)
	})

	// Good when you have two configurations, one for development and a different one for production use.
	app.Run(iris.Addr(":9090"), iris.WithConfiguration(c))
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

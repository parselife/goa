package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	xormCore "github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/core/errors"
	"goa/goa/core"
	"goa/goa/model"
	"time"
)

func Init(conf *core.SqlConf) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(conf.DriverName, conf.DbUrl)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	// 配置 engine 的sql 输出以及日志级别
	engine.ShowSQL(conf.ShowSql)
	engine.Logger().SetLevel(xormCore.LogLevel(conf.LogLevel))
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")

	// 同步表结构
	err = engine.Sync2(new(model.User), new(model.JobLog), new(model.JobProject),
		new(model.JobType), new(model.Organ), new(model.Message), new(model.MessageBody))
	if err != nil {
		return nil, err
	}
	return engine, nil
}

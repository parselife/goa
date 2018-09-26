package core

type AppConf struct {
	SqlConf
	AppInfo
}

// 数据库连接配置
type SqlConf struct {
	DriverName string
	DbUrl      string
	ShowSql    bool
	LogLevel   int64
}

// 应用信息配置
type AppInfo struct {
	Name   string
	Desc   string
	Author string
}

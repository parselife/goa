# goa
> 使用 `golang` 开发， 符合 `mvc` 架构的工时记录应用.

前端展示页面使用 `vuejs` 框架，项目地址: [goa-web](https://github.com/Yxf005/goa-web)

## 安装
`golang` 开发环境，推荐 **1.10+**
1. 获取源码

````shell
$ go get -u https://github.com/parselife/goa
````

2. 配置数据库连接

  打开 `/etc/app.tml`, 定位到

````toml
[Other.sqlConf]
    DriverName = "mysql"
    DbUrl = "root:qwer1234@tcp(127.0.0.1:3306)/goa_dev?charset=utf8"
    ShowSql = false
    LogLevel = 0
````

3. go build & run
 
  可选参数:
 - `-p 8000` 可改变端口
 - `-log2file` 将系统日志写入文件



## 技术路线
- [iris framework](https://github.com/kataras/iris)
- [xorm](https://github.com/go-xorm/xorm) 

## v0.1 实现功能
- [x] 权限控制
- [x] 月视图、周视图、日视图查看
- [x] 周汇报、月汇报功能
- [x] 消息提醒功能

## 下一版本计划
- [ ] 返回今天功能
- [ ] 工时提醒功能
- [ ] 读取禅道的项目与需求信息
- [ ] 完善用户体验

# goa
> 使用 `golang` 开发， 符合 `mvc` 特征的工时记录应用.

## 安装
`golang` 开发环境，推荐 **1.10+**
1. 获取源码

````shell
$ go get -u https://github.com/Yxf005/goa
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

## 下一版本计划
- [ ] 返回今天功能
- [ ] 工时提醒功能
- [ ] 读取禅道的项目与需求信息
- [ ] 完善用户体验
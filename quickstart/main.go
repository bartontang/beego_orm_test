package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "quickstart/routers"
)

var dbuser string = "root"       //数据库用户名
var dbpassword string = "123456" //数据库密码
var db string = "ormtest"        //数据库名字

//自动建表
func createTable() {
	name := "default"                          //数据库别名
	force := false                             //不强制建数据库
	verbose := true                            //打印建表过程
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		beego.Error(err)
	}
}

func init() {
	// 注册sqlite3 Driver
	// orm.RegisterDataBase("default", "sqlite3", "data.db")
	//
	//

	// 注册mysql Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 构造conn连接 用户名:密码@数据库地址+名称?字符集
	conn := dbuser + ":" + dbpassword + "@/" + db + "?charset=utf8"
	beego.Info(conn)
	//注册数据库连接
	orm.RegisterDataBase("default", "mysql", conn)
	createTable()
}

func main() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)
	// orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	//

	beego.Run()
}

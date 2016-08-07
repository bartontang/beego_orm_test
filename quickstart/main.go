package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	_ "quickstart/routers"
)

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
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	createTable()
}

func main() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)
	// orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	//

	beego.Run()
}

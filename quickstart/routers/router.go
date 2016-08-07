package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/test/:id:int", &controllers.TestController{}, "get:GetInfo")
	beego.Router("/test/:id:int/:name:string", &controllers.TestController{}, "get:UpdateInfo")
	beego.Router("/testmysql", &controllers.TestMySQLController{})
}

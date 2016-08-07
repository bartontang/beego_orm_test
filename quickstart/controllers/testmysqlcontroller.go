package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"quickstart/models"
)

type TestMySQLController struct {
	beego.Controller
}

func (c *TestMySQLController) Get() {
	c.Data["json"] = map[string]interface{}{"rc": 1,
		"msg":  "success",
		"data": false,
	}

	o := orm.NewOrm()
	o.Using("default")

	user, err := models.NewUser("barton")

	if err == nil {
		beego.Debug(o.Insert(user))
	}

	c.ServeJSON()
}

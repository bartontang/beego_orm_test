package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"quickstart/models"
	"strconv"
)

type TestController struct {
	beego.Controller
}

/**
 * Get请求
 * @param  每次请求对数据库插入一条新的user
 * @return {[type]}   [description]
 */
func (c *TestController) Get() {
	c.Data["json"] = map[string]interface{}{"rc": 1,
		"msg":  "success",
		"data": false,
	}

	o := orm.NewOrm()
	o.Using("default")

	// 插入一条User数据而且名字都赋值为barton
	user, err := models.NewUser("barton")
	if err == nil {
		beego.Debug(o.Insert(user))
	}

	c.ServeJSON()
}

/**
 * 获取User数据
 * @param  id唯一标识符（自增长）
 * @return {[type]}   [description]
 */
func (c *TestController) GetInfo() {
	id := c.Ctx.Input.Param(":id")
	beego.Trace("check id = " + id)
	o := orm.NewOrm()
	user := new(models.UserBT)

	user.Id, _ = strconv.ParseInt(id, 10, 64)
	err := o.Read(user)

	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"rc": 1,
			"msg":  "false",
			"data": "can not find id = " + id + "from table user",
		}
	} else {
		c.Data["json"] = map[string]interface{}{"rc": 1,
			"msg": "success",
			"data": map[string]interface{}{
				"userId":   user.Id,
				"userName": user.Name,
			},
		}
	}

	c.ServeJSON()
}

/**
 * 更新User数据
 * @param  id唯一标识符（自增长），name 想更新的新的名字
 * @return {[type]}   [description]
 */
func (c *TestController) UpdateInfo() {
	id := c.Ctx.Input.Param(":id")
	newName := c.Ctx.Input.Param(":name")
	beego.Trace("check id = " + id)
	beego.Trace("update name = " + newName)
	o := orm.NewOrm()
	user := new(models.UserBT)

	// 对id做int64的类型转换
	user.Id, _ = strconv.ParseInt(id, 10, 64)
	// 查找Id == id 的数据
	err1 := o.Read(user)

	// 查找出错并没找到对应的user数据
	if err1 != nil {
		beego.Error(err1)
		c.Data["json"] = map[string]interface{}{"rc": 1,
			"msg":  "false",
			"data": "can not find id = " + id + "from table user",
		}
	} else {
		user.Name = newName
		// 对查找的数据user进行name的更新
		_, err2 := o.Update(user, "name")
		if err2 != nil {
			// 更新失败
			beego.Error(err2)
			c.Data["json"] = map[string]interface{}{"rc": 1,
				"msg":  "false",
				"data": "update name id = " + id + "from table user",
			}
		} else {
			// 更新成功
			c.Data["json"] = map[string]interface{}{"rc": 1,
				"msg": "true",
				"data": map[string]interface{}{
					"userId":   user.Id,
					"userName": user.Name,
				},
			}
		}
	}

	c.ServeJSON()
}

func (c *TestController) Post() {
	account := c.GetString("account")
	password := c.GetString("password")
	beego.Info(account)

	c.Data["json"] = map[string]interface{}{"rc": 1,
		"msg": "success",
		"data": map[string]interface{}{
			"account":  account,
			"password": password,
		},
	}

	c.ServeJSON()
}

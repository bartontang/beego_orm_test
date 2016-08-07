package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserBT struct {
	Id   int64 `orm:"pk;auto"` //主键，自动增长
	Name string
}

func NewUser(name string) (*UserBT, error) {
	if name == "" {
		beego.Error("user name is emptry")
		return nil, fmt.Errorf("user name is emptry")
	}
	return &UserBT{Name: name}, nil
}

func init() {
	orm.RegisterModel(new(UserBT))

}

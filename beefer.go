package main

import (
	"github.com/astaxie/beego"
)

type BeeferController struct {
	beego.Controller
}

type User struct {
	Username string
}

func (c *BeeferController) Get() {

	user := User{Username: "腾"}
	c.Data["User"] = user

	c.TplNames = "beefer.tpl"
}

func main() {
	beego.Router("/", &BeeferController{})
	beego.Run(":3000")
}

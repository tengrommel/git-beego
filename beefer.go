package main

import (
	"github.com/astaxie/beego"
)

type BeeferController struct {
	beego.Controller
}

type UserController struct {
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

func (c *UserController) Login() {
	c.TplNames = "beefer.tpl"
}

func (c *UserController) Signup() {
	c.TplNames = "beefer.tpl"
}

func (c *UserController) Logout() {
	c.TplNames = "beefer.tpl"
}

func main() {
	beego.Router("/", &BeeferController{})
	beego.Router("/user/login", &UserController{}, "get:Login")
	beego.Router("/user/signup", &UserController{}, "get:Signup")
	beego.Router("/user/logout", &UserController{}, "post:Logout")
	beego.Run(":3000")
}

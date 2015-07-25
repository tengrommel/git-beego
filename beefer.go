package main

import (
	"github.com/astaxie/beego"
)

type BeeferController struct {
	beego.Controller
}

func (c *BeeferController) Get() {
	c.Ctx.WriteString("Hello Beego")
}

func main() {
	beego.Router("/", &BeeferController{})
	beego.Run(":3000")
}

package main

import (
	"github.com/astaxie/beego"
	"text/template"
)

type BeeferController struct {
	beego.Controller
}

type User struct {
	Username string
}

func (c *BeeferController) Get() {
	//c.Ctx.WriteString("Hello Beego")
	var tpl string = `
        <html>
            <head>
                <title>Beefer!</title>
            </head>
            <body>
                <strong>Hello, {{.User.Username}}</strong>
            </body>
        </html>
 `
	data := make(map[interface{}]interface{})
	user := User{Username: "腾"}
	data["User"] = user

	t := template.New("Beefer Template")
	t = template.Must(t.Parse(tpl))

	t.Execute(c.Ctx.ResponseWriter, data)
}

func main() {
	beego.Router("/", &BeeferController{})
	beego.Run(":3000")
}

package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) TestMe() {

	fmt.Println(c.Ctx.Input.Param(":cat1"))
	fmt.Println(c.Ctx.Input.Param(":cat2"))
	fmt.Println(c.Ctx.Input.Param(":cat3"))

	fmt.Println("URL: ", c.Ctx.Input.URL())
	fmt.Println("URI: ", c.Ctx.Input.URI())

	c.TplName = "index.tpl"
}

func (c *MainController) Get() {

	login := c.GetSession("login")
	id := c.GetSession("id")

	if login == nil {
		c.Redirect("/login", 302)
	}

	App.Logger.LogUrl(id.(int), "/", "")
	c.Redirect("/check/solaris", 302)

}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) LoginPage() {

	login := c.GetSession("login")
	id := c.GetSession("id")
	role := c.GetSession("root")

	if login != nil {
		App.Logger.LogUrl(id.(int), "/login", "get")

		if role == 1 {
			c.Redirect("/admin/statistics", 302)
		}
		c.Redirect("/check/solaris", 302)
	}

	App.Logger.LogUrl(-1, "/login", "get")
	c.TplName = "login.tpl"
}

func (c *LoginController) Authentication() {

	var response string

	name := c.Ctx.Request.FormValue("name")
	pass := c.Ctx.Request.FormValue("pass")

	user_id, role_id := App.Emloyers.CheckAuth(name, pass)

	if user_id != -1 {
		c.SetSession("login", name)
		c.SetSession("id", user_id)
		c.SetSession("role", role_id)
		response = "true"

		App.Logger.LogUrl(user_id, "/login", "Authentication")

	} else {
		response = "false"
		App.Logger.LogUrl(-1, "/login", "Authentication#"+name+"#"+pass)
	}

	c.Ctx.ResponseWriter.Write([]byte(response))
}

func (c *LoginController) Logout() {

	login := c.GetSession("login")
	id := c.GetSession("id")

	if login != nil {
		App.Logger.LogUrl(id.(int), "/logout", "get")
	} else {
		App.Logger.LogUrl(-1, "/logout", "get")
	}

	c.DelSession("login")
	c.DelSession("id")
	c.DelSession("role")

	c.Redirect("/", 302)
}

type TestController struct {
	beego.Controller
}

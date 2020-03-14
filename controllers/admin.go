package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Statistics() {

	access := false

	login := c.GetSession("login")
	//user_id := c.GetSession("id")
	role_id := c.GetSession("role")

	if login == nil {
		c.Redirect("/login", 302)
	}

	roles := App.Permissions.GetRoles(c.Ctx.Input.URL())

	for i := 0; i < len(roles); i++ {
		if role_id == roles[i] {
			access = true
		}
	}

	if !access {
		c.Redirect("/", 302)
	}

	if role_id == 1 {
		c.Data["Root"] = true
		c.Data["Managment"] = false
	} else if role_id == 2 {
		c.Data["Root"] = false
		c.Data["Managment"] = true
	} else {
		c.Data["Root"] = false
		c.Data["Managment"] = false
	}

	c.Layout = "layout/default.tpl"
	c.TplName = "/check/solaris"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Navigation"] = "element/navigation.tpl"

}

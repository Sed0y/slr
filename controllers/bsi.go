package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
)

type BSI struct {
	beego.Controller
}

func (b *BSI) Interface() {

	access := false

	login := b.GetSession("login")
	//user_id := c.GetSession("id")
	role_id := b.GetSession("role")
	//fmt.Println(login, " ", role_id)
	if login == nil {
		b.Redirect("/login", 302)
	}

	roles := App.Permissions.GetRoles(b.Ctx.Input.URL())

	for i := 0; i < len(roles); i++ {
		if role_id == roles[i] {
			access = true
		}
	}

	if !access && len(roles) != 0 {
		b.Redirect("/", 302)
	}

	if role_id == 1 {
		b.Data["Root"] = true
		b.Data["Managment"] = false
	} else if role_id == 2 {
		b.Data["Root"] = false
		b.Data["Managment"] = true
	} else {
		b.Data["Root"] = false
		b.Data["Managment"] = false
	}

	b.Layout = "layout/default.tpl"
	b.TplName = "checker/bsi.tpl"
	b.LayoutSections = make(map[string]string)
	b.LayoutSections["Navigation"] = "element/navigation.tpl"
}

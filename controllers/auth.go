package controllers

import (
	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Check() {

	login := c.GetSession("login")

	if login == nil {
		c.Ctx.ResponseWriter.Write([]byte("redirect"))
		return
	} else {
		c.Ctx.ResponseWriter.Write([]byte("auth"))
		return
	}

}

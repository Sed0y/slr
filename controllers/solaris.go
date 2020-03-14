package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type Solaris struct {
	beego.Controller
}

func (s *Solaris) Interface() {

	access := false

	login := s.GetSession("login")
	role_id := s.GetSession("role")

	if login == nil {
		s.Redirect("/login", 302)
	}

	roles := App.Permissions.GetRoles(s.Ctx.Input.URL())

	for i := 0; i < len(roles); i++ {
		if role_id == roles[i] {
			access = true
		}
	}

	if !access && len(roles) != 0 {
		s.Redirect("/", 302)
	}

	if role_id == 1 {
		s.Data["Root"] = true
		s.Data["Managment"] = false
	} else if role_id == 2 {
		s.Data["Root"] = false
		s.Data["Managment"] = true
	} else {
		s.Data["Root"] = false
		s.Data["Managment"] = false
	}

	s.Layout = "layout/default.tpl"
	s.TplName = "checker/solaris.tpl"
	s.LayoutSections = make(map[string]string)
	s.LayoutSections["Navigation"] = "element/navigation.tpl"

}

func (s *Solaris) GetPeopleFullRequest() {

	var uid_int int
	login := s.GetSession("login")
	user_id := s.GetSession("id")

	fmt.Println(login, "-", user_id)

	if login == nil {
		s.Ctx.ResponseWriter.Write([]byte("redirect"))
		return
	}

	if user_id != nil {
		uid_int = user_id.(int)
	}

	surname := s.Ctx.Request.FormValue("surname")
	name := s.Ctx.Request.FormValue("name")
	fathername := s.Ctx.Request.FormValue("fathername")
	bd_day := s.Ctx.Request.FormValue("bd_day")
	bd_month := s.Ctx.Request.FormValue("bd_month")
	bd_year := s.Ctx.Request.FormValue("bd_year")
	inn := s.Ctx.Request.FormValue("inn")

	params := "[{"
	params += "  \"s\":\"" + surname + "\","
	params += "  \"n\":\"" + name + "\","
	params += "  \"f\":\"" + fathername + "\","
	params += "  \"bdd\":\"" + bd_day + "\","
	params += "  \"bdm\":\"" + bd_month + "\","
	params += "  \"bdy\":\"" + bd_year + "\","
	params += "  \"i\":\"" + inn + "\""
	params += "}]"

	fmt.Println(params)

	response := App.Solaris.GetPeoplesFullRequest(surname, name, fathername, bd_day, bd_month, bd_year, inn)
	//fmt.Println(response)
	// Логировать ошибку выполнения запроса
	//App.Logger.LogSolaris(uid_int, "GetPeopleFullRequest", params)

	//fmt.Println(response.ParcedResponse)

	if len(response.ParcedResponse) > 0 {
		App.Logger.LogSolaris(uid_int, "GetPeopleFullReq", params)
		res := response.RenderResults(App.Solaris.Schemas)
		s.Ctx.ResponseWriter.Write([]byte(res))
	} else {
		if response.Response == "nothing" {
			s.Ctx.ResponseWriter.Write([]byte("nothing"))
			App.Logger.LogSolaris(uid_int, "nothing", params)
		} else {
			App.Logger.LogSolaris(uid_int, "error", "GetPeopleFullReq#"+response.Response)
			s.Ctx.ResponseWriter.Write([]byte("error###" + response.Response))
		}

	}

}

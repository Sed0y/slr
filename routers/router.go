package routers

import (
	"solaris/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})                              // log +
	beego.Router("/login", &controllers.LoginController{}, "get:LoginPage")       // log +
	beego.Router("/login", &controllers.LoginController{}, "post:Authentication") // log +
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")         // log +

	beego.Router("/auth/check", &controllers.AuthController{}, "post:Check")

	beego.Router("/admin", &controllers.AdminController{}, "get:Statistics")

	beego.Router("/admin/statistics", &controllers.AdminController{}, "get:Statistics")

	beego.Router("/charts", &controllers.ChartsController{}, "get:Charts")

	beego.Router("/api/?:cat1/?:cat2/?:cat3/filter", &controllers.MainController{}, "get:TestMe")

	//beego.Router("/solaris/people/byfio", &controllers.Solaris{}, "post:GetPeopleByFio")
	beego.Router("/solaris/people/full_request", &controllers.Solaris{}, "post:GetPeopleFullRequest")
	beego.Router("/check/solaris", &controllers.Solaris{}, "get:Interface") // log -

}

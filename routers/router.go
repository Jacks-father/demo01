package routers

import (
	"demo01/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//beego.Router("/user_register",&controllers.{})
}

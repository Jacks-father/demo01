package main

import (
	_ "demo01/routers"
	"github.com/astaxie/beego"
)

func main() {


	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}


package controllers

import (
	"demo01/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}
func (r *RegisterController)Post(){
	var user models.User
	err := r.ParseForm(&user)
	if err != nil{
		r.Ctx.WriteString("抱歉，解析数据错误，请重试！")
		return
	}
	_, err = user.SaveUser()
	if err != nil{
		r.Ctx.WriteString("抱歉，用户注册失败，请重试！")
		return
	}
	r.TplName = "log.html"
}
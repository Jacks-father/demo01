package controllers

import (
	"github.com/astaxie/beego"
	"DataCertProject/models"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post(){
	//1、解析请求数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		//返回错误信息给浏览器,提示用户
		r.Ctx.WriteString("抱歉，解析数据错误，请重试！")
		return
	}

	//2、保存用户信息到数据库
	_, err = user.SaveUser()

	//3、返回前端结果(成功跳登录页面，失败弹出错误信息）
	if err != nil {
		r.Ctx.WriteString("抱歉，用户注册失败，请重试!")
		return
	}
	//用户注册成功
	r.TplName = "login.html"

}

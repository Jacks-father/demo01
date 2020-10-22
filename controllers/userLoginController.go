package controllers

import (
	"github.com/astaxie/beego"
	"DataCertProject/models"
	"fmt"
)

type LoginController struct {
	beego.Controller
}

/**
 * 直接访问login.html页面的请求
 */
func (l *LoginController) Get(){
	//设置login.html为模板文件
	//tpl: template
	l.TplName = "login.html"
}

/**
 * 用户登录接口
 */
func (l  *LoginController) Post(){
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		//l.TplName = "error.html"
		l.Ctx.WriteString("抱歉，用户信息解析失败，请重试！")
		return
	}
	//fmt.Println(user.Phone,user.Password)
	//查询数据库的用户信息
	u, err := user.QueryUser()
	if err != nil {
		// sql: no rows in result set（集合）：结果集中无数据
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登录失败, 请重试!")
		return
	}
	//登录成功,跳转项目核心功能页面（home.html)
	l.Data["Phone"] = u.Phone
	l.TplName = "home.html"//{{.Phone}}
}

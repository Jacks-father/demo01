package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init()  {

	//1.读取conf配置信息
	config := beego.AppConfig

	dbDriver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	//2.组织链接数据库的字符串
	connUrl := dbUser + ":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	//fmt.Println(connUrl)
	// //fmt.Println(dbDriver)
	//3.链接数据库
	db, err := sql.Open(dbDriver,connUrl)
	if err != nil{
		fmt.Println(err.Error())
		panic("数据库连接错误，请检查配置")
	}
	//4.为全局变量赋值
	DB = db
}
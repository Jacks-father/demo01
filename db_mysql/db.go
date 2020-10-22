package db_mysql

import (
	"github.com/astaxie/beego"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

/**
 * 链接数据
 */
func ConnectDB(){
	//1、读取conf配置信息
	config := beego.AppConfig
	dbDriver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	//fmt.Println(dbDriver,dbUser,dbPassword,dbName)

	//2、组织链接数据库的字符串
	connUrl := dbUser +":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	//3、链接数据库
	db, err := sql.Open(dbDriver,connUrl)
	//4、获取数据库链接对象，处理链接结果
	if err != nil {// err不为nil，表示连接数据库时出现了错误, 程序就在此中断就可以，不用再执行了。
		//早发现，早解决
		panic("数据库连接错误，请检查配置")
	}

	Db = db//为全局的数据库操作对象赋值
}


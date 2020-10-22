package main

import (
	_ "DataCertProject/routers"
	"github.com/astaxie/beego"
	"DataCertProject/db_mysql"
	"DataCertProject/blockchain"
	"fmt"
	"DataCertProject/models"
	"encoding/json"
	"encoding/xml"
)

func main() {

	user1 := models.User{
		Id:       1,
		Phone:    "13167582311",
		Password: "abcdefg",
	}
	fmt.Println("内存中的数据User1:", user1)

	//json格式
	/**
	  * { "Id": 1, "phone":"131","Pasword":"fasdfa" }
	 */
	_, _ = json.Marshal(user1)
	xmlBytes, _ := xml.Marshal(user1)
	fmt.Println(string(xmlBytes))

	var user2 models.User
	xml.Unmarshal(xmlBytes, &user2)
	fmt.Println("反序列化的User2:", user2)
	return

	//1.生成第一个区块
	block := blockchain.NewBlock(0, []byte{}, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	fmt.Println(block)
	fmt.Printf("区块Hash值：%x\n", block.Hash)
	fmt.Printf("区块的nonce值:%d\n", block.Nonce)
	return
	//1、链接数据库
	db_mysql.ConnectDB()

	//2、静态资源路径设置
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")

	//3、允许
	beego.Run() //启动端口监听: 阻塞

}

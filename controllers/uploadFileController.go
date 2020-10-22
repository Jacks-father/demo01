package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"io"
	"bufio"
	"os"
	"DataCertProject/models"
	"time"
	"DataCertProject/util"
)

/**
 * 文件上传控制器
 */
type UploadFileController struct {
	beego.Controller
}

/**
 * 跳转新增认证文件的页面upload_file.html
 */
func (u *UploadFileController) Get() {
	phone := u.GetString("phone")
	u.Data["Phone"] = phone
	u.TplName = "home.html"
}

/**
 * 使用post方法处理文件的上传
 */
func (u *UploadFileController) Post() {
	//1、获取客户端上传的文件以及其他form表单的信息

	//标题
	fileTitle := u.Ctx.Request.PostFormValue("upload_title")
	phone := u.Ctx.Request.PostFormValue("phone")

	//文件
	file, header, err := u.GetFile("upload_file")
	if err != nil {
		u.Ctx.WriteString("抱歉，用户文件解析失败，请重试")
		return
	}
	//3、关闭文件
	defer file.Close()

	fmt.Println("自定义的文件标题:", fileTitle)
	fmt.Println("文件名称:", header.Filename)
	fmt.Println("文件的大小:", header.Size) //字节大小

	//2、将文件保存在本地的一个目录中
	//文件全路径： 路径 + 文件名 + "." + 扩展名
	//要的文件的路径
	uploadDir := "./static/img/" + header.Filename
	//文件权限：a+b+c
	//a:文件所有者拥有的权限，读4、写2、执行1
	//b:文件所有者所在的组的用户对文件拥有的权限，读4、写2、执行1
	//c:其他用户对文件拥有的权限，读4、写2、执行1
	//eg:某个文件m，其权限是985(错误)
	saveFile, err := os.OpenFile(uploadDir, os.O_RDWR|os.O_CREATE, 777)

	//创建一个writer: 用于向硬盘上写一个文件
	writer := bufio.NewWriter(saveFile)
	_, err = io.Copy(writer, file)
	if err != nil { //invalid argument
		fmt.Println(err.Error())
		u.Ctx.WriteString("抱歉，保存电子数据失败，请重试")
		return
	}
	defer saveFile.Close()

	//2、计算文件的hash
	hashFile, err := os.Open(uploadDir)
	defer hashFile.Close()
	hash, err := util.MD5HashReader(hashFile)

	//3、将上传的记录保存到数据库中
	record := models.UploadRecord{}
	record.FileName = header.Filename
	record.FileSize = header.Size
	record.FileTitle = fileTitle
	record.CertTime = time.Now().Unix() //毫秒数
	record.FileCert = hash
	record.Phone = phone //手机
	_, err = record.SaveRecord()
	if err != nil {
		u.Ctx.WriteString("抱歉，数据认证错误, 请重试!")
		return
	}

	//4、从数据库中读取phone用户对应的所有认证数据记录
	records, err := models.QueryRecordByPhone(phone)

	//5、根据文件保存结果，返回相应的提示信息或者页面跳转
	if err != nil {
		u.Ctx.WriteString("抱歉，获取认证数据失败, 请重试!")
		return
	}
	fmt.Println(records)
	u.Data["Records"] = records
	u.Data["Phone"] = phone
	u.TplName = "list_record.html"
}

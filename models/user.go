package models

import (
	"DataCertProject/db_mysql"
	"DataCertProject/util"
)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

/**
 * 保存用户信息的方法：保存用户信息到数据库中
 */
func (u User) SaveUser() (int64, error) {
	//1、密码脱敏处理
	u.Password = util.MD5HashString(u.Password)

	//2、执行数据库操作
	row, err := db_mysql.Db.Exec("insert into user(phone, password)"+
		" values(?,?) ", u.Phone, u.Password)
	if err != nil {
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

/**
 * 查询用户信息
 */
func (u User) QueryUser()(*User,error) {
	//1、密码脱敏处理
	u.Password = util.MD5HashString(u.Password)

	row := db_mysql.Db.QueryRow("select phone from user where  phone = ? and password = ?",
		u.Phone, u.Password)
	err := row.Scan(&u.Phone)
	if err != nil {
		return nil,err
	}
	return &u,nil
}

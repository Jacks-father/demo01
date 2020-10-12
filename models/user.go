package models

import (

	"crypto/md5"
	"demo01/db_mysql"
	"encoding/hex"
)

type  User struct {
	Id int          `form:"id"`
	Phone string      `form:"phone"`
	Password string       `form:"password"`
}
func (u User) SaveUser() (int64,error){
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes := md5Hash.Sum(nil)
	u.Password = hex.EncodeToString(passwordBytes)

	row,err:= db_mysql.DB.Exec("insert into user(phone,password)"+"values(?,?)",u.Phone,u.Password)
	if err !=nil{
		return -1,err
	}
	id,err :=row.RowsAffected()
	if err !=nil{
		return -1,err
	}
	return id,nil
}
func (u User)QueryUser()(*User,error){
	row := db_mysql.DB.QueryRow("select phone from user where phone = ? and password = ?"+
		u.Phone,u.Password)
	err :=row.Scan(&u.Phone)
	if err !=nil{
		return nil,err
	}
	return &u,nil
}
package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/astaxie/beego"
)
var Conn *gorm.DB
func init() {
	user :=beego.AppConfig.String("mysql_user")
	pwd :=beego.AppConfig.String("mysql_pwd")
	url :=beego.AppConfig.String("mysql_url")
	db_name :=beego.AppConfig.String("mysql_db_name")
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",user,pwd,url,db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil {
		panic("初始化mysql连接失败")
	}
	Conn =db
}

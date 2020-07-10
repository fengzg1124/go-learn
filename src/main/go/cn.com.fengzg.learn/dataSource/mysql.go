package dataSource

import (
	"github.com/jinzhu/gorm"
)

var(
	DB *gorm.DB
)

/*
	初始化连接数据库
 */
func InitMysql() (err error) {
	url := "root:root@tcp(localhost:3306)/go_learn?charset=utf8mb4&parseTime=True&loc=Local"
	DB,err = gorm.Open("mysql",url)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

/*
	关闭数据库
 */
func Close()  {
	DB.Close()
}
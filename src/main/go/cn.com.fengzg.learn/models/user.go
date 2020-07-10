package models

import "github.com/jinzhu/gorm"

/*
	用户表
 */
type User struct {
	gorm.Model
	UserName string
	password string
	createBy string
	updateBy string
	deleteBy string
	tel string
	age int8
}

/*
	表名
 */
func (User)TableName() string {
	return "tbl_user_info"
}
package dao

import (
	"learn/src/main/go/cn.com.fengzg.learn/dataSources"
	"learn/src/main/go/cn.com.fengzg.learn/models"
)

/*
	创建用户
*/
func CreateAUser(user *models.User) (err error) {
	err = dataSource.DB.Create(&user).Error
	return
}

/*
	根据id查询用户
*/
func GetAUser(id string) (user *models.User, err error) {
	user = new(models.User)
	err = dataSource.DB.Where("id=?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return
}

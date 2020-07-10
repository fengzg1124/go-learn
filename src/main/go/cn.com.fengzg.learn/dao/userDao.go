package dao

import (
	"learn/src/main/go/cn.com.fengzg.learn/dataSource"
	"learn/src/main/go/cn.com.fengzg.learn/models"
)

func CreateAUser(user *models.User) (err error) {
	err = dataSource.DB.Create(&user).Error
	return
}

func getAUser(id string) (user *models.User, err error)  {
	user = new(models.User)
	err = dataSource.DB.Where("id=?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return
}

package logic

import (
	"fmt"
	"learn/src/main/go/cn.com.fengzg.learn/dao"
	"learn/src/main/go/cn.com.fengzg.learn/models"
)

/*
	创建用户
*/
func CreateAUser(user *models.User) {
	err := dao.CreateAUser(user)
	if err != nil {
		fmt.Println("create user err:" + err.Error())
	}
}

/*
	根据id查询用户
*/
func GetAUser(id string) (user *models.User) {
	user, err := dao.GetAUser(id)
	if err != nil {
		fmt.Println("get user err:" + err.Error())
		return nil
	}
	return
}

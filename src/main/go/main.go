package main

import (
	"fmt"
	"learn/src/main/go/cn.com.fengzg.learn/dataSources"
	"learn/src/main/go/cn.com.fengzg.learn/models"
	"learn/src/main/go/cn.com.fengzg.learn/routers"
)

func main() {

	//连接数据库

	if err := dataSources.InitMysql(); err != nil {
		panic(err)
	}

	//退出程序关闭数据库
	defer dataSources.Close()

	//模型绑定
	dataSources.DB.AutoMigrate(&models.User{})

	//加载多个路由配置
	routers.Include(routers.UserRouters)

	//路由初始化
	r := routers.Init()

	if err := r.Run(":8081"); err != nil {
		fmt.Println("startup service failed, err:" + err.Error())
	}
}

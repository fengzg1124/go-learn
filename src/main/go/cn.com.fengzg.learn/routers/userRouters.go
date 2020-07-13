package routers

import (
	"github.com/gin-gonic/gin"
	"learn/src/main/go/cn.com.fengzg.learn/controller"
)

func UserRouters(e *gin.Engine) {
	userGroup := e.Group("user")

	{
		userGroup.POST("/add", controller.CreateAUser)
		userGroup.GET("/get", controller.GetAUser)
	}
}

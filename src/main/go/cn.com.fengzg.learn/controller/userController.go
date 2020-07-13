package controller

import (
	"github.com/gin-gonic/gin"
	"learn/src/main/go/cn.com.fengzg.learn/logic"
	"learn/src/main/go/cn.com.fengzg.learn/models"
	"net/http"
)

func CreateAUser(c *gin.Context) {

	var user = new(models.User)
	c.BindJSON(user)
	logic.CreateAUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "",
		"data":    "",
	})

}

func GetAUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "无效的id",
			"data":    "",
		})
		return
	}
	user := logic.GetAUser(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "",
		"data":    user,
	})
}

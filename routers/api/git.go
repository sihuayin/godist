package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sihuayin/godist/models"
	"github.com/sihuayin/godist/pkg/components"
)

func GetGitBranch(c *gin.Context) {

	projectId := c.DefaultQuery("projectId", "0")
	pid, _ := strconv.Atoi(projectId)

	if pid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "参数错误",
			"data": nil,
		})
		return
	}
	pro, _ := models.GetProjectById(pid)
	gitCom := components.NewGitComponent(&pro)
	list, _ := gitCom.GetBranchList()

	// userInterface, ok := c.Get("User")
	// if !ok {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "ss首保"})
	// 	return
	// }

	// user := userInterface.(*models.User)

	// pro.Name = pro.Name + "- CPOY"
	// pro.Id = 0
	// pro.UserId = uint(user.ID)
	// newPro, _ := pro.Create()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": list,
	})
}

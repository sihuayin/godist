package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sihuayin/godist/models"
)

func CopyPostConf(c *gin.Context) {
	userInterface, ok := c.Get("User")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ss首保"})
		return
	}

	user := userInterface.(*models.User)
	projectId := c.DefaultQuery("projectId", "0")
	pid, _ := strconv.Atoi(projectId)
	pro, _ := models.GetProjectById(pid)

	pro.Name = pro.Name + "- CPOY"
	pro.Id = 0
	pro.UserId = uint(user.ID)
	newPro, _ := pro.Create()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": newPro,
	})
}

func DeletePostConf(c *gin.Context) {
	projectId := c.DefaultQuery("projectId", "0")
	pid, _ := strconv.Atoi(projectId)
	err := models.DeleteProject(pid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除失败",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "删除成功",
		"data": nil,
	})
	return
}
func PostConfInfo(c *gin.Context) {
	projectId := c.DefaultQuery("projectId", "0")
	pid, _ := strconv.Atoi(projectId)
	project, _ := models.GetProjectById(pid)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": project,
	})
	return
}

func PostConfOfMine(c *gin.Context) {
	where := ""
	userInterface, ok := c.Get("User")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ss首保"})
		return
	}

	user := userInterface.(*models.User)
	if user.Role == 10 {
		where += "AND `level`=2 "
	} else if user.Role == 20 {
		where += "and  id in (SELECT project_id FROM `group` WHERE `group`.user_id=" + strconv.Itoa(int(user.ID)) + " )  "
	}
	projects, _ := models.FindProjects(where, 0, 10)
	data := map[string]interface{}{"total": 10, "currentPage": 1, "table_data": projects}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}
func PostConfList(c *gin.Context) {
	// page := c.DefaultQuery("page", "0")
	// start := 0
	// length := c.DefaultQuery("length", "200000")
	// if page > 0 {
	// 	start = (page - 1) * length
	// }
	// selectInfo := c.DefaultQuery("select_info", "")
	// where := ""
	// if selectInfo != "" {
	// 	where = "  and(`name` LIKE '%" + selectInfo + "%' )"
	// }
	projects, _ := models.FindProjects("", 0, 10)
	data := map[string]interface{}{"total": 10, "currentPage": 1, "table_data": projects}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})

	// o.Raw("SELECT *, (SELECT realname FROM `user` WHERE `user`.id=project.user_id LIMIT 1) as realname FROM `project`  WHERE 1=1 "+where+" ORDER BY id LIMIT ?,?", start, length).Values(&projects)
	// var count []orm.Params
	// total := 0
	// o.Raw("SELECT count(id) FROM `project` WHERE 1=1 " + where).Values(&count)
	// if len(count) > 0 {
	// 	total = common.GetInt(count[0]["count(id)"])
	// }
	// c.SetJson(0, map[string]interface{}{"total": total, "currentPage": page, "table_data": projects}, "")

	// return
}
func PostConfSave(c *gin.Context) {
	var project models.Project
	var err error
	if err = c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if project.Id != 0 {
		err = project.Save()
	} else {
		_, err = project.Create()
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": project,
	})
}

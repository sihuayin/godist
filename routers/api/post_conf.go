package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sihuayin/godist/models"
)

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
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": projects,
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

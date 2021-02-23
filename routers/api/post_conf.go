package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sihuayin/godist/models"
)

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

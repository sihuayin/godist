package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sihuayin/godist/pkg/e"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// valid := validation.Validation{}
	// a := auth{Username: username, Password: password}
	// ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	data["username"] = username
	data["password"] = password

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

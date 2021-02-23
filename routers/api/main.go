package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GoHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tpl", gin.H{
		"title": "Main website11",
	})
}

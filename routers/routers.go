package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sihuayin/godist/middleware/header"
	"github.com/sihuayin/godist/pkg/setting"
	api "github.com/sihuayin/godist/routers/api"
	// v1 "github.com/sihuayin/go-gin-demo/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(header.AuthString())
	gin.SetMode(setting.RunMode)
	r.GET("/", api.GoHome)

	r.POST("/login", api.GetAuth)
	r.POST("/logout", api.AuthLogout)
	r.POST("/changePasswd", api.AuthChangePWD)

	r.GET("/api/get/conf/list", api.PostConfList)
	r.POST("/api/post/conf/save", api.PostConfSave)
	r.GET("/api/get/conf/get", api.PostConfInfo)
	r.GET("/api/get/conf/del", api.DeletePostConf)
	r.GET("/api/get/conf/copy", api.CopyPostConf)
	r.GET("/api/get/conf/mylist", api.PostConfOfMine)

	r.GET("/api/get/git/branch", api.GetGitBranch)

	// apiv1 := r.Group("/api/v1")
	// {

	// }
	return r
}

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

	r.POST("/api/post/conf/save", api.PostConfSave)

	// apiv1 := r.Group("/api/v1")
	// {

	// }
	return r
}

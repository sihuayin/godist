package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/sihuayin/godist/models"
	"github.com/sihuayin/godist/pkg/setting"
	"github.com/sihuayin/godist/routers"
)

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncdb()
			os.Exit(0)
		}
	}
}

func main() {
	initArgs()
	router := routers.InitRouter()
	router.LoadHTMLGlob("views/*")
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

package main

// import "github.com/gin-gonic/gin"

import (
	"fmt"

	"github.com/sihuayin/godist/pkg/setting"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
	fmt.Println(setting.HTTPPort)
}

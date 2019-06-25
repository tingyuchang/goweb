package routers

import (
	"github.com/gin-gonic/gin"

	"goweb/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	// regist routers

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "test",
		})
	})

	return r
}

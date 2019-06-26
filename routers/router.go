package routers

import (
	"github.com/gin-gonic/gin"

	"goweb/pkg/setting"
	"goweb/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		// tag
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// artivle
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/article/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/article/:id", v1.EditArticle)
		apiv1.DELETE("/article/:id", v1.DeleteArticle)
	}

	return r
}

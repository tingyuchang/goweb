package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/astaxie/beego/validation"
	"github.com/Unknwon/com"

	"goweb/models"
	"goweb/pkg/e"
	"goweb/pkg/setting"
	"goweb/pkg/util"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	state := -1

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddTag(c *gin.Context) {
}

func EditTag(c *gin.Context) {
}

func DeleteTag(c *gin.Context) {
}

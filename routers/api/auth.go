package api

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"goweb/models"
	"goweb/pkg/e"
	"goweb/pkg/util"
)

type Auth struct {
	Username string `valid:"Required; MaxSize(50)" form:"username"`
	Password string `valid:"Required; MaxSize(50)" form:"password"`
}

func GetAuth(c *gin.Context) {
	var auth Auth
	c.ShouldBind(&auth)

	log.Println(auth)

	valid := validation.Validation{}
	ok, _ := valid.Valid(&auth)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(auth.Username, auth.Password)
		if isExist {
			token, err := util.GenerateToken(auth.Username, auth.Password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

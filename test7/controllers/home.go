package controllers

import (
	"test7/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
		return
	}

	c.Data["Topics"] = topics
	c.TplName = "home.html"
}

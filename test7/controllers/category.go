package controllers

import (
	"test7/models"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {

	op := this.Input().Get("op")
	var err error

	//beego.Info(op)

	switch op {
	case "add":
		category := this.Input().Get("category")
		if len(category) == 0 {
			break
		}

		err = models.AddCategory(category)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err = models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return
	}

	this.TplName = "category.html"
	this.Data["IsCategory"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}

package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type PostController struct {
	beego.Controller
}

type PostData struct {
	name string
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (this *PostController) Post() {
	fmt.Println("hello")
	//获取form-data的数据
	name := this.GetString("name")
	fmt.Println(name)
	//获取body中的json
	var data map[string]interface{}
	fmt.Println(this.Ctx.Input.RequestBody)
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	fmt.Println(data["name"])
}

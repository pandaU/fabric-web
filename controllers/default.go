package controllers

import (
	"encoding/json"
	"fabric-web/models"
	"fabric-web/request"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type MainController struct {
	beego.Controller
}

/*func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.bbcom"
	c.TplName = "index.html"
}*/
func (c *MainController) Post() {
	DB :=models.Conn
	req :=request.UserResqust{}
	err :=json.Unmarshal(c.Ctx.Input.RequestBody,&req)
	if err != nil {
		c.Data["json"]=models.RespError("请求参数错误")
		c.ServeJSON()
		return
	}
	valid := validation.Validation{}
	b, err := valid.Valid(&req)
	if err != nil || !b {
		c.Data["json"]=models.RespError("请求参数格式错误")
		c.ServeJSON()
		return
	}
	result :=&models.UserInfo{}
	DB.Model(result).Where(req).Select("truename,id,phone").Find(result)
	c.Data["json"]=models.RespSuccess(result)
	c.ServeJSON()
}

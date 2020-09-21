package controllers

import (
	"encoding/json"
	"fabric-web/models"
	"fabric-web/request"
	"fabric-web/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var redisClient utils.RedisUtil
	redisClient.Set("name","pandaU",30*time.Second)
	c.Data["json"]=models.RespSuccess(redisClient.Get("name"))
	c.ServeJSON()

}
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

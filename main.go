package main

import (
	_ "fabric-web/models"
	_ "fabric-web/routers"
	_ "fabric-web/utils"
	"github.com/astaxie/beego"
)

func main() {
	//beego.LoadAppConfig("ini","conf/redis.conf")
	beego.Run()
}


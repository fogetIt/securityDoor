package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/admin"
	_ "github.com/astaxie/beego/session/redis"
	_ "securityDoor/models"
	_ "securityDoor/filters"
	_ "securityDoor/routers"
)


func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		admin.Run()
	}
	beego.Run()
}

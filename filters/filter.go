package filters


import (
	"strings"
	//"reflect"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"securityDoor/models"
	"securityDoor/utils"
	"github.com/astaxie/beego/orm"
)


/*
TODO  ip 白名单拦截
 */
var verifyAdmin = func(ctx *context.Context)  {
	if Ip := ctx.Input.IP(); Ip != "127.0.0.1" {
		if orm.NewOrm().QueryTable("ip_white").
			Filter("Ip", Ip).Exist() {
			ctx.Abort(401, "ip error")
		}
	}
}


/*
TODO  接口流量统计
 */
func verifySso(ctx *context.Context) {
	appName := ctx.Request.Header["appName"]
	var jsonStr string
	if len(appName) == 0 {
		jsonStr = utils.CreateMessage(0, "appName error")
	} else {
		al := models.AppLog{}
		field := strings.Split(ctx.Input.URL(), "/")[2]
		if b := al.UpdateCallCount(appName[0], field); !b {
			jsonStr = utils.CreateMessage(0, "call interface error")
		}
	}
	if jsonStr != "" {
		ctx.Output.JSON(jsonStr, true, true)
	}
}


func init() {
	// TODO  跨域
	beego.InsertFilter("*",
		beego.BeforeRouter,
		cors.Allow(&cors.Options{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			AllowCredentials: true,
		}))

	beego.InsertFilter("/public/*", beego.BeforeExec, verifyAdmin)
	beego.InsertFilter("/sso/*", beego.BeforeExec, verifySso)
}

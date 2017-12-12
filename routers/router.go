// @APIVersion 1.0.0
// @Title voteDoor API
// @Description vote sso application, routers start with /sso
// @Contact 2271404280@qq.com
// @TermsOfServiceUrl
// @License
// @LicenseUrl
package routers


import (
	"github.com/astaxie/beego"
	"securityDoor/controllers"
)


func init() {
	/*
	TODO  beego 自动文档
	目前只支持 Namespace + NSInclude
	 */
	ns := beego.NewNamespace("/sso",
		beego.NSNamespace("/code",
			beego.NSInclude(
				&controllers.CodeController{})),

		beego.NSNamespace("/session",
			beego.NSInclude(
				&controllers.SessionController{})),

		beego.NSNamespace("/token",
			beego.NSInclude(
				&controllers.TokenController{})),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{})),
	)
	beego.AddNamespace(ns)
}
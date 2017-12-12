package controllers

import (
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"securityDoor/models"
	"securityDoor/utils"
)


type TokenController struct {
	beego.Controller
}


// @Title get user token
// @Description login API
// @Param Email query string false Email
// @Param Mobile query int false Mobile
// @Param Password query string true Password
// @Success 200 {json}
// @Failure
// @router / [get]
func (this *TokenController) Get() {
	Email := this.GetString("Email")
	Mobile, _ := this.GetUint16("Mobile")
	Password := this.GetString("Password")

	var jsonStr string
	if Password == "" {
		jsonStr = utils.CreateMessage(0, "password error")
	}else if Email == "" && Mobile == 0  {
		jsonStr = utils.CreateMessage(0, "mobile and email are all empty")
	} else if !orm.NewOrm().
		QueryTable("user").
		Filter("Email", Email).
		Filter("Mobile", Mobile).
		Exclude("Status", 0).Exist() {
		jsonStr = utils.CreateMessage(0, "user not exist")
	} else {
		var us models.User
		err := orm.NewOrm().
			QueryTable("user").
			Filter("Mobile", Mobile).
			Filter("Email", Email).One(&us)
		if err == nil {
			Pwd := us.Pwd
			if !utils.VerifyPwd(Password, Pwd) {
				jsonStr = utils.CreateMessage(0, "password does not match")
			} else {
				addr := this.Ctx.Input.IP()

				user := models.User{
					Mobile: Mobile,
					Email: Email}
				user.LastLoginIp = addr
				user.Update("LastLoginIp", "LastLoginAt", "ModifyAt")

				var tokenMap = make(map[string]interface{})
				token := utils.GenerateToken(strconv.Itoa(int(us.UserId)))
				tokenMap["token"] = token
				jsonStr = utils.CreateMessage(1, "login successful", tokenMap)
			}
		} else {
			jsonStr = utils.CreateMessage(0, "database error")
		}
	}
	this.Data["json"] = jsonStr
	this.ServeJSON()
}
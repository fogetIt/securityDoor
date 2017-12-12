package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"securityDoor/utils"
	"securityDoor/models"
)


type SessionController struct {
	beego.Controller
}


func (this *SessionController) validateParams() uint {
	Token := this.GetString("Token")
	UserId, _ := this.GetUint16("UserId")
	if Token == "" && UserId == 0 {
		return 0
	} else if UserId == 0 {
		b := utils.VerifyToken(Token)
		if b != 0 {
			return b
		} else {
			return 0
		}
	} else {
		return uint(UserId)
	}
}


// @Title get user session
// @Description get user info API
// @Param Token query string false Token
// @Param UserId query int false UserId
// @Success 200 {json}
// @Failure
// @router / [get]
func (this *SessionController) Get() {
	var jsonStr string
	Id := this.validateParams()
	if Id == 0 {
		jsonStr = utils.CreateMessage(0, "UserId and Token error")
	} else {
		userSession := this.GetSession(Id)
		if userSession == nil {
			var us models.User
			var UserMap map[string]interface{}
			err := orm.NewOrm().QueryTable(new(models.User)).
				Filter("UserId", Id).One(&us)
			if err == nil {
				UserMap = make(map[string]interface{})
				UserMap["Email"] = &us.Email
				UserMap["Status"] = &us.Status
				UserMap["Mobile"] = &us.Mobile
				UserMap["UserId"] = &us.UserId
				UserMap["CreateIp"] = &us.CreateIp
				UserMap["UserName"] = &us.UserName
				UserMap["CreateAt"] = &us.CreateAt
				UserMap["ModifyAt"] = &us.ModifyAt
				UserMap["LastLoginIp"] = &us.LastLoginIp
				UserMap["LastLoginAt"] = &us.LastLoginAt
				this.SetSession(Id, UserMap)
			} else {
				jsonStr = utils.CreateMessage(0, "get user info error")
			}
		}
	}
	if jsonStr == "" {
		jsonStr = utils.CreateMessage(1, "get user info successful",
			this.GetSession(Id).(map[string]interface{}))
	}
	this.Data["json"] = jsonStr
	this.ServeJSON()
}


// @Title delete user session
// @Description logout API
// @Param Token body string false Token
// @Param UserId body int false UserId
// @Success 200 {json}
// @Failure
// @router / [delete]
func (this *SessionController) Delete() {
	var jsonStr string
	Id := this.validateParams()
	if Id == 0 {
		jsonStr = utils.CreateMessage(0, "UserId and Token error")
	} else {
		this.DelSession(Id)
		jsonStr = utils.CreateMessage(0, "logout successful")
	}
	this.Data["json"] = jsonStr
	this.ServeJSON()
}

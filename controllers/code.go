package controllers


import (
	"fmt"
	"time"
	"math/rand"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"securityDoor/utils"
	"securityDoor/models"
)


const (
	ErrorLimit = 5
	AttemptsLimit = 10
	CodeIntervalTimes = 1 * 60
	CodeErrorFreezeTimes = 60 * 60
)


type CodeController struct {
	beego.Controller
}


func (this *CodeController) validateUser(Mobile uint16, Email string, UserStatus uint8) string {
	if UserStatus == 0{
		return "UserStatus error"
	} else if Mobile == 0 && Email == "" {
		return "mobile and email are all empty"
	} else {
		if UserStatus == 0 {
			if orm.NewOrm().
				QueryTable("user").
				Filter("Email", Email).
				Filter("Mobile", Mobile).
				Exclude("Status", 0).Exist() {
				return "mobile and email are registered"
			}
		} else {
			if !orm.NewOrm().
				QueryTable("user").
				Filter("Email", Email).
				Filter("Mobile", Mobile).
				Filter("Status", UserStatus).Exist() {
				return "mobile and email are not registered"
			}
		}
	}
	return ""
}


func (this *CodeController) sendVerificationCode(Email string, Mobile uint16, Code int) string {
	var sessionKey interface{}
	var errorTimes uint8 = 0
	var attemptsTimes uint8 = 0
	if Email == "" {
		sessionKey = Email
	} else {
		sessionKey = Mobile
	}

	userSession := this.GetSession(sessionKey).(map[interface{}]interface{})
	if userSession != nil {
		freeze := userSession["freeze"].(bool)
		timestamp := userSession["timestamp"].(int64)
		errorTimes = userSession["errorTimes"].(uint8)
		attemptsTimes = userSession["attemptsTimes"].(uint8)
		if freeze {
			if attemptsTimes + 1 > AttemptsLimit || errorTimes > ErrorLimit {
				userSession["code"] = 0
				userSession["freeze"] = true
				userSession["timestamp"] = time.Now().Unix()
				userSession["errorTimes"] = 0
				userSession["attemptsTimes"] = 0
				return "too many attempts, please try after 1 hour"
			} else if time.Now().Unix() - timestamp < CodeIntervalTimes {
				return "please try after 1 minute"
			}
		} else {
			if time.Now().Unix() - timestamp < CodeErrorFreezeTimes {
				return fmt.Sprintf("you %s is frozen, please try after 1 hour", sessionKey)
			}
		}
	} else {
		this.SetSession(sessionKey, make(map[string]interface{}))
	}
	var result bool
	if Email == "" {
		result = utils.SendEmail(Code)
	} else {
		result = utils.SendMessage(Code)
	}
	if result {
		userSession["code"] = Code
		userSession["freeze"] = false
		userSession["timestamp"] = time.Now().Unix()
		userSession["errorTimes"] = errorTimes
		userSession["attemptsTimes"] = attemptsTimes + 1
		return ""
	} else {
		return "mobile number or email address error"
	}
}


// @Title get verification code
// @Description send verification code to user when register or modify password API
// @Param Email query string false Email
// @Param Mobile query int false Mobile
// @Param UserStatus query int true UserStatus
// @Success 200 {json} {status: 0/1, msg: ""}
// @Failure
// @router / [get]
func (this *CodeController) Get() {
	Email := this.GetString("Email")
	Mobile, _ := this.GetUint16("Mobile")
	UserStatus, _ := this.GetUint8("UserStatus")

	Code := rand.Intn(999999)
	fmt.Println(Code)

	var jsonStr string
	if errorText := this.validateUser(Mobile, Email, UserStatus); errorText != "" {
		jsonStr = utils.CreateMessage(0, errorText)
	} else if errorText := this.sendVerificationCode(Email, Mobile, Code); errorText != "" {
		jsonStr = utils.CreateMessage(0, errorText)
	} else {
		user := models.User{Mobile: Mobile, Email: Email, Status: 0}
		// 返回参数：是否新创建， Id ，错误
		if created, id, err := orm.NewOrm().
			ReadOrCreate(&user, "Mobile", "Email", "Status"); err == nil {
			if created {
				fmt.Println("New Insert an object. Id:", id)
			} else {
				fmt.Println("Get an object. Id:", id)
			}
			jsonStr = utils.CreateMessage(1, "send verification code successful")
		} else {
			jsonStr = utils.CreateMessage(0, "send verification code successful, but database error")
		}
	}
	this.Data["json"] = jsonStr
	this.ServeJSON()
}

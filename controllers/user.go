package controllers


import (
	"fmt"
	"time"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"securityDoor/models"
	"securityDoor/utils"
)

const (
	CodeExpiredDTime = 3 * 60
)


type UserController struct {
	beego.Controller
}


func (this *UserController) verifyCode(Code int, sessionKey string) string {
	//userSession := context.Context{}.Input.CruSession.Get(sessionKey).(map[string]interface{})
	userSession := this.GetSession(sessionKey).(map[string]interface{})
	if userSession == nil {
		return "verification code is not in session"
	} else if Code != userSession["Code"] {
		errorTimes := userSession["errorTimes"].(int)
		userSession["errorTimes"] = errorTimes + 1
		return "verification code not match"
	} else if time.Now().Unix() - userSession["timestamp"].(int64) > CodeExpiredDTime {
		return "verification code is timeout"
	}
	return ""
}


func (this *UserController) validateUser(Email string, Mobile uint16, UserStatus uint8) string {
		if Email == "" && Mobile == 0  {
			return "mobile and email are all empty"
		}
		if UserStatus == 0 {
			if orm.NewOrm().
				QueryTable("user").
				Filter("Email", Email).
				Filter("Mobile", Mobile).
				Exclude("Status", 0).Exist() {
				return "mobile and email are registered"
			} else if !orm.NewOrm().
				QueryTable("user").
				Filter("Email", Email).
				Filter("Mobile", Mobile).
				Filter("Status", 0).Exist() {
				return "mobile and email are not validated"
			}
		} else {
			if !orm.NewOrm().
				QueryTable("user").
				Filter("Email", Email).
				Filter("Mobile", Mobile).
				Filter("Status", UserStatus).Exist() {
				return "mobile and email are not validated"
			}
		}
		return ""
}


// @Title register / modify password
// @Description register / modify password API
// @Param Email body string false Email
// @Param Mobile body int false Mobile
// @Param Password body string true Password
// @Param UserName body string false UserName
// @Param Code body int true Code
// @Param UserStatus body int true UserStatus
// @Success 200 {json}
// @Failure
// @router / [post]
func (this *UserController) Post() {
	Email := this.GetString("Email")
	Password := this.GetString("Password")
	UserName := this.GetString("UserName")
	Code, _ := this.GetInt("Code")
	Mobile, _ := this.GetUint16("Mobile")
	UserStatus, _ := this.GetUint8("UserStatus")

	var jsonStr, sessionKey string
	if Code == 0 || Password == "" {
		jsonStr = utils.CreateMessage(0, "Code and Password are necessary")
	} else {
		if errorText := this.validateUser(Email, Mobile, UserStatus); errorText != "" {
			jsonStr = utils.CreateMessage(0, errorText)
		} else {
			if Email != "" {
				sessionKey = Email
			} else {
				sessionKey = strconv.Itoa(int(Mobile))
			}
			if errorText := this.verifyCode(Code, sessionKey); errorText != "" {
				jsonStr = utils.CreateMessage(0, errorText)
			} else {
				this.DelSession(sessionKey)
				user := models.User{
					Mobile: Mobile,
					Email: Email,
					Status: UserStatus}
				user.Pwd = Password
				user.Update("Pwd")

				if UserStatus == 0 {
					addr := this.Ctx.Input.IP()

					user.CreateIp = addr
					user.UserName = UserName
					user.Status = 1
					user.Update("CreateIp", "UserName", "Status", "CreateAt")
				}

				var us models.User
				var UserType string
				err := orm.NewOrm().QueryTable(new(models.User)).
					Filter("Mobile", Mobile).
					Filter("Email", Email).One(&us)
				if UserStatus == 0 {
					UserType = "modify password"
				} else {
					UserType = "register"
				}
				if err == nil {
					var UserIdMap map[string]interface{}      // TODO  声明
					UserIdMap = make(map[string]interface{})  // TODO  初始化
					UserIdMap["UserId"] = us.UserId
					jsonStr = utils.CreateMessage(1,
						fmt.Sprintf("%s successful", UserType), UserIdMap)
				} else {
					jsonStr = utils.CreateMessage(0,
						fmt.Sprintf("%s successful, but database error", UserType))
				}
			}
		}
	}
	this.Data["json"] = jsonStr
	this.ServeJSON()
}

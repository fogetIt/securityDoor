package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)


type IpWhite struct {
	BaseModel
	Id          uint      `orm:"pk;auto;column(id);unique"`
	Ip          string    `orm:"size(20);default('')"`
	Description string    `orm:"size(20);default('')"`
	Creator     string    `orm:"size(20);unique"`
	CreateAt    time.Time `orm:"auto_now_add;type(datetime)"`
	ModifyAt    time.Time `orm:"auto_now;type(datetime)"`
}


type App struct {
	BaseModel
	AppId    uint      `orm:"pk;auto;column(id);unique"`
	AppName  string    `orm:"size(100);unique"`
	Creator  string    `orm:"size(20);unique"`
	CreateAt time.Time `orm:"auto_now_add;type(datetime)"`
	ModifyAt time.Time `orm:"auto_now;type(datetime)"`
}


type AppLog struct {
	BaseModel
	Id               uint      `orm:"pk;auto;column(id);unique"`
	AppId            uint      `orm:"rel(fk)"`
	Date             time.Time `orm:"auto_now;type(date)"`
	CodeCallCount    int64     `orm:"default(0)"`
	SessionCallCount int64     `orm:"default(0)"`
	TokenCallCount   int64     `orm:"default(0)"`
	UserCallCount    int64     `orm:"default(0)"`
}


func (this *AppLog) UpdateCallCount(appName string, field string) bool {
	var app App
	var appLog AppLog
	if err := orm.NewOrm().QueryTable("App").
		Filter("AppName", appName).One(&app); err == nil {
		AppId := app.AppId
		today := time.Now().AddDate(0, 0, 0)
		al := AppLog{Date: today, AppId: AppId}
		if o, b := al.Read(); !b {
			al.Date = today
			al.AppId = AppId
			_, err := o.Insert(&al)
			if err != nil {
				return false
			}
		}
		if err := orm.NewOrm().QueryTable("AppLog").
			Filter("AppId", AppId).Filter("Date", today).
			One(&appLog); err == nil{
			if field == "CodeCallCount" {
				al.CodeCallCount = appLog.CodeCallCount + 1
			} else if field == "SessionCallCount" {
				al.SessionCallCount = appLog.SessionCallCount + 1
			} else if field == "UserCallCount" {
				al.UserCallCount = appLog.UserCallCount + 1
			} else if field == "TokenCallCount" {
				al.TokenCallCount = appLog.TokenCallCount + 1
			}
			al.Update()
		}
	}
	return false
}

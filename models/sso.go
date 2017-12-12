package models

import (
	"time"
)


/*
TODO
默认的表名、字段名规则，使用驼峰转蛇形
更新单条数据库记录，如果只传入单个参数，会将其他没传入的字段置空
 */
type User struct {
	UserId      uint      `orm:"pk;auto;column(id);unique"`
	Status      uint8     `orm:"default(0)"`
	Mobile      uint16    `orm:"default(0)"`
	Email       string    `orm:"size(50);default('')"`
	UserName    string    `orm:"size(20);null"`
	CreateIp    string    `orm:"size(20);default('')"`
	LastLoginIp string    `orm:"size(20);default('')"`
	Pwd         string    `orm:"size(120);null"`
	CreateAt    time.Time `orm:"auto_now_add;type(datetime)"`
	ModifyAt    time.Time `orm:"auto_now;type(datetime)"`
	LastLoginAt time.Time `orm:"type(datetime);null"`
}

package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


func init() {
	//beego.LoadAppConfig("ini", "conf/app.conf")
	dbLink := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s",
		beego.AppConfig.String("db_user"),
		beego.AppConfig.String("db_pass"),
		beego.AppConfig.String("db_host"),
		beego.AppConfig.String("db_port"),
		beego.AppConfig.String("db_name"),
		beego.AppConfig.String("db_charset"))

	maxIdle := 30
	maxConn := 30

	orm.RegisterDriver(beego.AppConfig.String("db_type"), orm.DRMySQL)

	// set default database
	orm.RegisterDataBase("default", "mysql", dbLink, maxIdle, maxConn)

	orm.Debug = true

	// register model
	orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("default", false, true)
}

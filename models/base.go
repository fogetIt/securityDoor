package models

import (
	"github.com/astaxie/beego/orm"
)


type ModelMethods interface {
	Read()
	Update()
}


type BaseModel struct {}


func (this *BaseModel) Read() (o orm.Ormer, b bool) {
	o = orm.NewOrm()
	b = false
	if o.Read(&this) == nil {
		b = true
		return
	}
	return
}


func (this *BaseModel) Update(cols ...string) int64 {
	o, b := this.Read()
	if b {
		if num, err := o.Update(&this, cols...); err == nil {
			return num
		}
	}
	return 0
}

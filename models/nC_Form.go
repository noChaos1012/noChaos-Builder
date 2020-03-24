package models

import (
	"time"
)

//数据结构
type NC_Form struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Name      string        //名称
	Fields    []NC_Field    `gorm:"FOREIGNKEY:FormId`     //字段
	Relations []NC_Relation `gorm:"FOREIGNKEY:FromFormId` //关联
}

//数据结构 — 字段
type NC_Field struct {
	ID            uint `gorm:"primary_key"`
	FormId        uint
	Name          string
	Type          string
	AutoIncrement bool   //自增
	Index         bool   //索引
	Unique        bool   //唯一
	NotNull       bool   //非空
	Default       string //默认
	Size          int    //长度
}

//关联
type NC_Relation struct {
	Type       string //关联类型、一对多、多对多
	FromFormId uint   //来自对象
	ToFormId   uint   //关联对象
}

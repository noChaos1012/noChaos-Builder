package models

import (
	"github.com/jinzhu/gorm"
)

//构建程序代码
type NC_Build struct {
	gorm.Model
	ServletId          uint `gorm:"not null"`
	Version            string
	VersionDescription string
}

//构建程序源码
func (p NC_Build) Build() {

}

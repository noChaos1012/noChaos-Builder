package models

import "github.com/jinzhu/gorm"

//服务
type NC_Servlet struct {
	gorm.Model
	Name        string
	Description string
}

//打包
type NC_Package struct {
}

//部署
type NC_Deploy struct {
}

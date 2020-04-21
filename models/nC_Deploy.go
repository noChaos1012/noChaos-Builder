package models

import "github.com/jinzhu/gorm"

//部署
type NC_Deploy struct {
	gorm.Model
	ServletId          uint `gorm:"not null"`
	Version            string
	VersionDescription string

	IP   string
	Port string
}

//部署
type NC_DeployDownload struct {
	gorm.Model
	ServletId          uint `gorm:"not null"`
	Version            string
	VersionDescription string
	IP                 string
	Port               string
}

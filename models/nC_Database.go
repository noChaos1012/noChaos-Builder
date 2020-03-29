package models

import "github.com/jinzhu/gorm"

//数据库
type NC_Database struct {
	gorm.Model
	ServletId uint
	Port      string
	Host      string
	UserName  string
	Password  string
	Charset   string `gorm:"varchar(255)"`
}

//创建数据库
func (servlet NC_Servlet) CreateDB() {
	NCDB.Debug().First(&servlet)

	NCDB.Exec("Create Database If Not Exists " +
		servlet.GetName() +
		" Character Set UTF8 Collate utf8_general_ci ")
}

package models

import (
	"fmt"
)

//数据库
type NC_Database struct {
	ServletId uint
	Port      string
	Host      string
	UserName  string
	Password  string
	DataBase  string
	Charset   string `gorm:"varchar(255)"`
}

// 获取连接代码
func (db NC_Database) GetConnectionCode() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", db.UserName, db.Password, db.Host, db.Port, db.DataBase, db.Charset)
}

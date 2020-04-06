package models

//数据库
type NC_Database struct {
	ID        uint `gorm:"primary_key"`
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

//数据库
type NC_FTP struct {
	ID        uint `gorm:"primary_key"`
	ServletId uint
	Port      string
	Host      string
	UserName  string
	Password  string
	Charset   string `gorm:"varchar(255)"`
}

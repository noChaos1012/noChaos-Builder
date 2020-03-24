package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var NCDB *gorm.DB

func init() {
	var err error

	NCDB, err = gorm.Open("mysql", "root:wyw940524@/noChaos-server?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		NCDB.DB().SetMaxIdleConns(10)
		NCDB.DB().SetMaxOpenConns(100)
	}
	//defer NCDB.Close()
	NCDB.AutoMigrate(&NC_Form{}, &NC_Field{}, &NC_Relation{}, &NC_Servlet{})
}

//分页结构体
type NC_Page struct {
	PageNo     int
	PageSize   int
	TotalCount int
	List       interface{}
}

func PageUtil(count int, pageNo int, pageSize int, list interface{}) NC_Page {

	return NC_Page{
		PageNo:     pageNo,
		PageSize:   pageSize,
		TotalCount: count,
		List:       list,
	}
}

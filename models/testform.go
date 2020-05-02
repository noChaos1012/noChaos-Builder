package models

type ShilibiaoN struct {
	Changwenben  string  `json:"Changwenben" gorm:"column:Changwenben;type:mediumtext;not null"`
	Zhengshu32   int32   `json:"Zhengshu32" gorm:"column:Zhengshu32;type:int;AUTO_INCREMENT;default:0"`
	Zhengshu64   int64   `json:"Zhengshu64" gorm:"column:Zhengshu64;type:bigint;index;unique"`
	Xiaoshu32    float32 `json:"Xiaoshu32" gorm:"column:Xiaoshu32;type:float(6,2)"`
	Xiaoshu64    float64 `json:"Xiaoshu64" gorm:"column:Xiaoshu64;type:double(5,2)"`
	Buer         bool    `json:"Buer" gorm:"column:Buer;type:int"`
	Shijian      string  `json:"Shijian" gorm:"column:Shijian;type:time"`
	Riqi         string  `json:"Riqi" gorm:"column:Riqi;type:date"`
	Riqishijian  string  `json:"Riqishijian" gorm:"column:Riqishijian;type:datetime"`
	Shijianchuo  string  `json:"Shijianchuo" gorm:"column:Shijianchuo;type:timestamp"`
	Changwenben2 string  `json:"Changwenben2" gorm:"column:Changwenben2;type:varchar(255);default:'默认文本2'"`
	Id           int32   `json:"Id" gorm:"column:Id;type:int;primary_key;AUTO_INCREMENT;default:0"`
}

//Email   string  `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
//`json:"field" xml:"demo"`
//长文本 不可以加unique index 默认值

package consts

import "time"

// barcodeLookup
var (
	NumberFormatGds = 5 //间隔个数
	TimeUnitGds     = 7 //间隔停顿时间 单位秒
	SentPauseGds    = 6 //每条数据停顿时间 单位秒
)

//中国商品信息平台
var (
	ExcelFile         = "C:/Users/Administrator/Desktop/klasdiof5.xlsx" //excel文件地址
	NumberFormat      = 10                                              //间隔个数
	TimeUnit          = 4                                               //间隔停顿时间 单位秒
	SentPause         = 1                                               //每条数据停顿时间 单位秒
	BarcodeMax   int8 = 15
	BarcodeMin   int8 = 8
)

const (
	Reptile_Success string = "数据爬取成功"
	Reptile_Fail    string = "数据爬取失败"
)

type AlcoholDepotReptile struct {
	Id           int
	Barcode      string    `gorm:"not null;unique;type:varchar(15)"` //条码
	Formats      string    `gorm:"type:varchar(5)"`                  //条码格式
	CnName       string    `gorm:"type:varchar(100)"`                //中文酒名
	EnName       string    `gorm:"type:varchar(255)"`                //外文酒名
	Brand        string    `gorm:"type:varchar(255)"`                //品牌
	Category     string    `gorm:"type:varchar(255)"`                //类别
	Manufacturer string    `gorm:"type:varchar(255)"`                //制造商
	Description  string    `gorm:"type:varchar(5000)"`               //产品描述
	Features     string    `gorm:"type:varchar(5000)"`               //酒品特征
	Length       string    `gorm:"type:varchar(20)"`                 //长
	Width        string    `gorm:"type:varchar(20)"`                 //宽
	Height       string    `gorm:"type:varchar(20)"`                 //高
	SizeUnit     string    `gorm:"type:varchar(10)"`                 //长度单位
	Weight       string    `gorm:"type:varchar(20)"`                 //重
	Quality      string    `gorm:"type:varchar(10)"`                 //重量单位
	Image        string    `gorm:"type:varchar(5000)"`               //图片地址
	Standard     string    `gorm:"type:varchar(100)"`                //规格型号
	Web          string    `grom:"type:varchar(2000)"`               //数据来源网址
	GmtCreate    time.Time //创建时间
	GmtModified  time.Time //更新时间
	Type         int8      `gorm:"default:0"` //商品状态（0:未爬取,1:已爬取,2:爬取失败）
	IsAlcohol    int8      `gorm:"default:1"` //判断是否为酒款(1:酒款,2:非酒款)
}

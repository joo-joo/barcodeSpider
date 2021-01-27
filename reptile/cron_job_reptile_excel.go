package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"spider/constant"
	"strconv"
	"time"
)

//barcodeLookup
func main() {
	cookie := "_ga=GA1.2.1848819408.1606961635; __gads=ID=742075db66b3e83d:T=1606997600:S=ALNI_MbtNAEJAREj66L6a-2YvjlUAyQ4Zw; __tawkuuid=e::barcodelookup.com::l0K3ZLDS7WMXYKzr8jPtLS8zJrVtLN4JRVIxtOPwFIH8h18mA45Dr1aHIekjI1vc::2; __cfduid=d1e931491f1a7ddd0a9be36e44f0c98941609722948; FCCDCF=[[\"AKsRol8SKru4vZEo3lKyEIAebzyu0skSIkslQtZKVhhD3oNGbODJcC-wippmHgIs1HYuUr6m8ABex5rh2SekZ_uiMOE0zlBof7iXFM9WY34oi6Z1lHgteGyadXTu_DvOA5RF5Y2roktwB5piE2KyLSS5uYI8qShiQQ==\"],null,[\"[[],[],[],[],null,null,true]\",1611114931193]]; cf_chl_prog=a19; cf_clearance=a8cd5d9002cd372efaad2fb3cad334004f59b1f1-1611540583-0-250; bl_session=4gngcatrqka8m2p53v03ero3blkm1k50; __cflb=04dToRCegghj9KSg7BvgXvyGCMdwUyfh5QaBSXD18R"
	var err error
	f, err := excelize.OpenFile(consts.ExcelFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetCols("alcohol_depot_reptile")
	for i, row := range rows[1] {
		if i > 0 {
			if int8(len(row)) < consts.BarcodeMin || int8(len(row)) > consts.BarcodeMax {
				continue
			}
			//针对网站设置停顿时间
			if i%consts.NumberFormat == 0 {
				time.Sleep(time.Duration(consts.TimeUnit) * time.Second)
			}
			if rows[21][i] == "1" {
				continue
			}
			if rows[21][i] == "2" {
				continue
			}
			if rows[21][i] == "3" {
				continue
			}
			if rows[21][i] == "4" {
				continue
			}
			barinfo, info, err := ReptileV2(cookie, row)
			if err != nil {
				fmt.Println(info, err)
				continue
			}
			if info == "403" || info == "something wrong" {
				fmt.Println(info + " 当前爬取条码为：" + row + "在第" + strconv.Itoa(i) + "列")
				break
			}
			f.SetCellValue("alcohol_depot_reptile", "C"+strconv.Itoa(i+1), barinfo.Formats)
			f.SetCellValue("alcohol_depot_reptile", "D"+strconv.Itoa(i+1), barinfo.CnName)
			f.SetCellValue("alcohol_depot_reptile", "E"+strconv.Itoa(i+1), barinfo.EnName)
			f.SetCellValue("alcohol_depot_reptile", "F"+strconv.Itoa(i+1), barinfo.Brand)
			f.SetCellValue("alcohol_depot_reptile", "G"+strconv.Itoa(i+1), barinfo.Category)
			f.SetCellValue("alcohol_depot_reptile", "H"+strconv.Itoa(i+1), barinfo.Manufacturer)
			f.SetCellValue("alcohol_depot_reptile", "I"+strconv.Itoa(i+1), barinfo.Description)
			f.SetCellValue("alcohol_depot_reptile", "J"+strconv.Itoa(i+1), barinfo.Features)
			f.SetCellValue("alcohol_depot_reptile", "K"+strconv.Itoa(i+1), barinfo.Length)
			f.SetCellValue("alcohol_depot_reptile", "L"+strconv.Itoa(i+1), barinfo.Width)
			f.SetCellValue("alcohol_depot_reptile", "M"+strconv.Itoa(i+1), barinfo.Height)
			f.SetCellValue("alcohol_depot_reptile", "N"+strconv.Itoa(i+1), barinfo.SizeUnit)
			f.SetCellValue("alcohol_depot_reptile", "O"+strconv.Itoa(i+1), barinfo.Weight)
			f.SetCellValue("alcohol_depot_reptile", "P"+strconv.Itoa(i+1), barinfo.Quality)
			f.SetCellValue("alcohol_depot_reptile", "Q"+strconv.Itoa(i+1), barinfo.Image)
			f.SetCellValue("alcohol_depot_reptile", "R"+strconv.Itoa(i+1), barinfo.Web)
			f.SetCellValue("alcohol_depot_reptile", "S"+strconv.Itoa(i+1), barinfo.Standard)
			f.SetCellValue("alcohol_depot_reptile", "U"+strconv.Itoa(i+1), barinfo.GmtModified.Format("2006/01/02 15:04:05"))
			f.SetCellValue("alcohol_depot_reptile", "V"+strconv.Itoa(i+1), barinfo.Type)
			f.SetActiveSheet(0)
			if err := f.Save(); err != nil {
				fmt.Println(err)
			}
			time.Sleep(time.Duration(consts.SentPause) * time.Second)
		}
	}
	fmt.Println("success")
}

package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"spider/constant"
	"strconv"
	"time"
)

//中国商品信息平台
func main() {
	//1.从待爬取excel文件中选出条码数据
	//任务参数		param.ExecutorParams
	cookie := "ASP.NET_SessionId=pywuzca24ywn2vaqaut1cven"
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
			if i%consts.NumberFormatGds == 0 {
				time.Sleep(time.Duration(consts.TimeUnitGds) * time.Second)
			}
			if rows[21][i] == "0" {
				continue
			}
			if rows[21][i] == "1" {
				continue
			}
			if rows[21][i] == "3" {
				continue
			}
			if rows[21][i] == "4" {
				continue
			}
			barinfo, info, err := ReptileAnccnetV2(cookie, row)
			if err != nil {
				fmt.Println(info, err)
				continue
			}
			if info == "403" || info == "something wrong" {
				fmt.Println(info + " 当前爬取条码为：" + row + "在第" + strconv.Itoa(i) + "列")
				fmt.Println(info)
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
			time.Sleep(time.Duration(consts.SentPauseGds) * time.Second)
		}
	}
	fmt.Println("success")

}

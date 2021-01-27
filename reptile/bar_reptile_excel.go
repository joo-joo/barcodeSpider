package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"spider/constant"
	"time"
)

//爬虫根据条码爬取信息
//使用excel的方式存储数据
func ReptileV2(cookie string, barCode string) (barinfo *consts.AlcoholDepotReptile, info string, err error) {
	var flag string
	b := &consts.AlcoholDepotReptile{}
	c := colly.NewCollector(
		colly.MaxDepth(1),
		colly.Debugger(&debug.LogDebugger{}))
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36"
	//////配置两个代理
	//rp, err := proxy.RoundRobinProxySwitcher("http://secondtransfer.moguproxy.com:9001", "http://transfer.moguproxy.com:9001")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//c.SetProxyFunc(rp)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "www.barcodelookup.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Cache-Control", "max-age=0")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Headers.Set("Cookie", cookie)
		r.Headers.Set("Referer", "https://www.barcodelookup.com/search") //关键头 如果没有 则返回 错误
		r.Headers.Set("Sec-Fetch-Site", "same-origin")
		r.Headers.Set("Sec-Fetch-Mode", "navigate")
		r.Headers.Set("Sec-Fetch-User", "?1")
		r.Headers.Set("Sec-Fetch-Dest", "document")
		r.Headers.Set("Accept-Encoding", "")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,und;q=0.7")
	})

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	//在OnResponse之后被调用，如果收到的内容是HTML
	c.OnHTML("#editProductForm > div.product-fields", func(e *colly.HTMLElement) {
		Image := make([]string, 6)
		b.Features = ""
		b.Image = ""
		b.Formats = e.ChildText(".barcode>div>.input-group-addon>label")
		b.Barcode = e.ChildAttr(".barcode>div>input", "value")
		//目标网站条码补0导致两个条码位数不同，统一按照数据库条码为准
		if b.Barcode != barCode {
			b.Barcode = barCode
		}
		//暂且默认网站不会根据条码出来无关商品，可以采集同酒多码的数据
		//code := e.ChildAttr(".barcode>div>input", "value")
		////目标网站条码补0导致两个条码位数不同，统一按照数据库条码为准 （使用正则防止抓错数据）
		//match1, _ := regexp.MatchString("^0*"+barCode, code)
		//if !match1 {
		//	b.Type = 2
		//	return
		//}
		//b.Barcode = barCode
		b.Category = e.ChildAttr("#scrollable-dropdown-menu > input", "value")
		b.EnName = e.ChildAttr(".productName>div>input", "value")
		b.Manufacturer = e.ChildAttr(".manufacturer>div>input", "value")
		b.Brand = e.ChildAttr(".brand>div>input", "value")
		b.Length = e.ChildAttr(".dimensions>div>div>input[name='length']", "value")
		b.Width = e.ChildAttr(".dimensions>div>div>input[name='width'] ", "value")
		b.Height = e.ChildAttr(".dimensions>div>div>input[name='height']", "value")
		b.SizeUnit = e.ChildText(".dimensions>div>div>select>option:nth-child(1)")
		b.Weight = e.ChildAttr(".weight>div>div>input", "value")
		b.Quality = e.ChildText(".weight>div>div>select>option:nth-child(1)")
		b.Description = e.ChildText(".description>div>textarea")
		e.ForEach("div.form-group.feature>div>input", func(i int, f *colly.HTMLElement) {
			b.Features += f.Attr("value")
		})
		Image[0] = e.ChildAttr("div:nth-child(33)>div>div>div>img", "src")
		Image[1] = e.ChildAttr("div:nth-child(34)>div>div>div>img", "src")
		Image[2] = e.ChildAttr("div:nth-child(35)>div>div>div>img", "src")
		Image[3] = e.ChildAttr("div:nth-child(36)>div>div>div>img", "src")
		Image[4] = e.ChildAttr("div:nth-child(37)>div>div>div>img", "src")
		Image[5] = e.ChildAttr("div:nth-child(38)>div>div>div>img", "src")
		for j, v := range Image {
			if v != "" {
				if j > 0 {
					b.Image += ";"
				}
				b.Image += v
			}
		}
		b.GmtModified = time.Now()
		if len(b.EnName) > 0 {
			b.Type = 1
		}
	})

	//在发起请求前被调用
	c.OnRequest(func(r *colly.Request) {
		b.Barcode = barCode
		b.Web = r.URL.String()
		b.GmtModified = time.Now()
		b.Type = 2
	})

	//请求过程中如果发生错误被调用
	c.OnError(func(_ *colly.Response, err error) {
		//报错的话type不要修改下次继续爬取
		b.Type = 0
		if err.Error() == "Forbidden" {
			fmt.Println("Forbidden 403 Please update cookie:", err)
			flag = "403"
		} else {
			fmt.Println("Something went wrong:", err)
			flag = "something wrong"
		}
	})

	//爬取地址
	c.Visit("https://www.barcodelookup.com/" + barCode)
	c.Wait()

	if flag != "" {
		return nil, flag, err
	}

	if len(b.EnName) == 0 {
		//if err = dao.UpdateBarCode(b); err != nil {
		//	log.Error(err.Error())
		//	return consts.Reptile_Db_Fail, err
		//}
		return b, consts.Reptile_Fail, err

	} else {
		//if err = dao.UpdateBarCode(b); err != nil {
		//	log.Error(err.Error())
		//	return consts.Reptile_Db_Fail, err
		//}
		return b, consts.Reptile_Success, err
	}

}

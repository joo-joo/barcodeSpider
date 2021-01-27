package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"regexp"
	"spider/constant"
	"time"
)

//爬虫根据条码爬取信息
//使用excel的方式存储数据
func ReptileAnccnetV2(cookie string, barCode string) (barinfo *consts.AlcoholDepotReptile, info string, err error) {
	var flag string
	b := &consts.AlcoholDepotReptile{}
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"), colly.MaxDepth(1), colly.Debugger(&debug.LogDebugger{}))
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "search.anccnet.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("Proxy-Authorization", "Basic OHo0aGFxYkNXNm9rZzVxWDpMM3pDMFRkWk1Cc2xJd2Va")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Headers.Set("Cookie", cookie)
		r.Headers.Set("Referer", "http://www.gds.org.cn/") //关键头 如果没有 则返回 错误
		r.Headers.Set("Accept-Encoding", "")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,und;q=0.7")
	})
	//文章列表
	c.OnHTML(".mainly", func(e *colly.HTMLElement) {
		match1, _ := regexp.MatchString("人机识别验证", e.ChildText("#Label4"))
		match2, _ := regexp.MatchString("(下市)|(暂无相关信息)", e.ChildText("#Label1"))
		if match1 {
			b.Type = 2
			fmt.Println("cookie失效了")
			flag = "cookie失效了"
			return
		}
		if match2 {
			fmt.Println("无匹配")
			return
		}
		code := e.ChildText("#results > li:nth-child(1) > div > dl.p-info > dd:nth-child(2) > a")
		match4, _ := regexp.MatchString("^0*"+barCode, code)
		if !match4 {
			b.Type = 3
			return
		}
		photo := e.ChildAttr("#repList_ctl00_herl > img", "src")
		match3, _ := regexp.MatchString("http", photo)
		if match3 {
			b.Image = photo
		}
		b.CnName = e.ChildText("#results > li > div > dl.p-info > dd:nth-child(6)")
		b.Standard = e.ChildText("#results > li > div > dl.p-info > dd:nth-child(8)")
		b.Description = e.ChildText("#results > li > div > dl.p-info > dd:nth-child(10)")
		b.Brand = e.ChildText("#results > li > div > dl.p-supplier > dd:nth-child(2)")
		b.Manufacturer = e.ChildText("#results > li > div > dl.p-supplier > dd:nth-child(4) > a")
		if len(b.CnName) > 0 {
			b.Type = 1
		}

	})

	//在发起请求前被调用
	c.OnRequest(func(r *colly.Request) {
		b.Barcode = barCode
		b.Web = r.URL.String()
		b.GmtModified = time.Now()
		b.Type = 3
	})

	//请求过程中如果发生错误被调用
	c.OnError(func(_ *colly.Response, err error) {
		//报错的话type不要修改下次继续爬取
		b.Type = 2
		if err.Error() == "Forbidden" {
			fmt.Println("Forbidden 403 Please update cookie:", err)
			flag = "403"
			return
		}
		if err.Error() == "Unauthorized" {
			fmt.Println("try again:", err)
			flag = "too fast"
			return
		}
		fmt.Println("Something went wrong:", err)
		flag = "something wrong"
		return
	})

	c.Visit("http://search.anccnet.com/searchResult2.aspx?keyword=" + barCode)

	if flag != "" {
		return nil, flag, err
	}
	if len(b.CnName) == 0 {
		return b, consts.Reptile_Fail, err
	} else {
		return b, consts.Reptile_Success, err
	}

}

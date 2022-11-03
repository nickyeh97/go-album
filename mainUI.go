package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type IndexData struct {
	Title   string
	Content string
}

func main() {
	// func main
	url := "https://zh.pngtree.com/free-animal-vectors"
	// url := "https://ithelp.ithome.com.tw/users/20125192/ironman/3155"
	c := colly.NewCollector() // 在colly中使用 Collector 這類物件 來做事情

	// 在發起請求之前，可以預先對Header的參數進行設定
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("1")
		fmt.Println("Visiting", r.URL)
		// iT邦幫忙需要寫這一段 User-Agent才給爬
		// r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	// 如果在請求的時候發生錯誤
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("2")
		fmt.Println("Visiting error", err)
	})

	c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
		fmt.Println("4")
		// fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println("5")
	})

	c.OnXML("//footer", func(e *colly.XMLElement) {
		fmt.Println("6")
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("7")
	})
	// 抓img Class
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("src"))
	})

	c.Visit(url) // Visit 要放最後

	ginWeb()
}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

var session string = "buvid3=A971D746-B561-0F1E-F07B-849387FB18BA77761infoc; b_nut=1717650877; _uuid=8928B859-5814-784E-101E8-8798F74D5410A76662infoc; buvid4=9E54CA59-BEBB-03F1-B9BB-188B293E277D78450-024060605-LTAan2OkZP%2BqLVxDEsEg0A%3D%3D; enable_web_push=DISABLE; home_feed_column=5; DedeUserID=201438388; DedeUserID__ckMd5=a18f046fbe4d67d5; CURRENT_FNVAL=4048; rpdid=|(um)~ukuukk0J'u~u~Yl|ml|; header_theme_version=CLOSE; buvid_fp_plain=undefined; fingerprint=ad91b0a460ba9cdc64fad3a37f8b8240; LIVE_BUVID=AUTO6617195514583163; PVID=1; buvid_fp=ad91b0a460ba9cdc64fad3a37f8b8240; CURRENT_QUALITY=80; bp_t_offset_201438388=956825444039000064; bsource=search_baidu; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjI2NDk2MjUsImlhdCI6MTcyMjM5MDM2NSwicGx0IjotMX0.xQPnaEvAVB8bdM95dG-AaC35SKhrVyu3AjbsHmdp1Es; bili_ticket_expires=1722649565; SESSDATA=ee92ff49%2C1737942426%2C39c5a%2A72CjA1F44tPHELZ1vjMoRIbctoMyOJh-OZ-wdx0jG0mm8G1TP_D_Ld7Jmjddp-yvsx5IUSVnVweEdPbElLVm0xRVdnbWE2dmVVOFlsUWRfaV9IUlp3SjhXNWE3MUU1TVBUb0RtekduMnB6ZjZySE05U3hOd1Jqb0NkeDhGa1JxRUFOdl9yR0VMWG1BIIEC; bili_jct=b1ed02e48e4969957162e28c9fa15185; b_lsid=8D53252A_191076ED271; browser_resolution=1488-714; bmg_af_switch=1; bmg_src_def_domain=i1.hdslb.com; sid=dh6g2aky"

func main() {
	//创建采集器对象
	c := colly.NewCollector(
		colly.UserAgent("360Spider"),
	)

	// 在访问页面之前执行的回调函数
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		r.Headers.Set("accept-encoding", "gzip, deflate, br, zstd")
		r.Headers.Set("Cookie", session)
		fmt.Println("Visiting", r.URL.String())
	})

	// 在访问页面之后执行的回调函数
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL.String())
	})

	// 在访问页面时发生错误时执行的回调函数
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		println(e.Text)
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach(".feed-card", func(i int, el *colly.HTMLElement) {
			el.ForEach("img", func(i int, h *colly.HTMLElement) {
				url := h.Attr("src")
				urlArr := strings.Split(url, "@")
				if url != "" {
					fmt.Println("发现img，地址为：", urlArr[0])
					resp, err := http.Get("http:" + urlArr[0])
					if err != nil {
						log.Fatal(err)
					}
					defer resp.Body.Close()
					name := h.Attr("alt")

					filename := "./img/" + name + ".jpg"
					file, err := os.Create(filename)
					defer file.Close()
					if err != nil {
						fmt.Println("create file fail:", err)
					} else {
						io.Copy(file, resp.Body)
						fmt.Println(name, " is saved")
					}
				}
			})
		})
	})

	// 对visit的线程数做限制，visit可以同时运行多个
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		//Delay:      5 * time.Second,
	})

	c.Visit("http://bilibili.com")

}

func baidu() {
	// NewCollector(options ...func(*Collector)) *Collector
	// 声明初始化NewCollector对象时可以指定Agent，连接递归深度，URL过滤以及domain限制等
	c := colly.NewCollector(
		//colly.AllowedDomains("news.baidu.com"),
		colly.UserAgent("Opera/9.80 (Windows NT 6.1; U; zh-cn) Presto/2.9.168 Version/11.50"),
	)

	// 发出请求时附的回调
	c.OnRequest(func(r *colly.Request) {
		// Request头部设定
		r.Headers.Set("Host", "baidu.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", "")
		r.Headers.Set("Referer", "http://www.baidu.com")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN, zh;q=0.9")

		fmt.Println("Visiting", r.URL)
	})

	// 对响应的HTML元素处理
	c.OnHTML("title", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		fmt.Println("title:", e.Text)
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		// <div class="hotnews" alog-group="focustop-hotnews"> 下所有的a解析
		e.ForEach(".hotnews a", func(i int, el *colly.HTMLElement) {
			band := el.Attr("href")
			title := el.Text
			fmt.Printf("新闻 %d : %s - %s\n", i, title, band)
			// e.Request.Visit(band)
		})
	})

	// 发现并访问下一个连接
	//c.OnHTML(`.next a[href]`, func(e *colly.HTMLElement) {
	//	e.Request.Visit(e.Attr("href"))
	//})

	// extract status code
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response received", r.StatusCode)
		// 设置context
		// fmt.Println(r.Ctx.Get("url"))
	})

	// 对visit的线程数做限制，visit可以同时运行多个
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		//Delay:      5 * time.Second,
	})

	c.Visit("http://news.baidu.com")
}

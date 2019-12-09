// http://top.baidu.com
package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// New 新闻
type New struct {
	IncidentTitle     string  `json:"incidentTitle"`
	LongTitle         string  `json:"longTitle"`
	LabelNames        string  `json:"labelNames"`
	RatioHotTopCustom float64 `json:"ratioHotTopCustom"`
	RatioHotDay       float64 `json:"ratioHotDay"`
	AreaType          int     `json:"areaType"`
	Province          string  `json:"province"`
	City              string  `json:"city"`
	Origin            string  `json:"origin"`
}

// Event 事件
type Event struct {
	Name       string `json:"name"`
	Summary    string `json:"summary"`
	LabelNames string `json:"labelNames"`
	AreaType   int    `json:"areaType"`
	Province   string `json:"province"`
	City       string `json:"city"`
}

// GetNews 获取新闻
func GetNews() []New {
	news := []New{}

	data := url.Values{}
	data.Set("sort", "2")
	data.Set("labels", "")
	data.Set("areaType", "")
	data.Set("province", "")
	data.Set("city", "")
	data.Set("startTime", GetDate(-1)+" 00:00:00")
	data.Set("endTime", GetDate(0)+" 00:00:00")
	data.Set("page", "1")
	data.Set("pageSize", "1000")
	data.Set("showTag", "1")
	data.Set("labelShowTag", "1")

	fmt.Println(data)

	body, _ := PostForm("http://m.wrd.cn/getHotEventList", data)

	type Resp struct {
		List []New `json:"list"`
	}

	var resp Resp
	json.Unmarshal([]byte(body), &resp)

	for i, new := range resp.List {
		if len(news) == 10 {
			break
		}

		fmt.Println(i, new.IncidentTitle, new.LongTitle, new.RatioHotDay, new.RatioHotTopCustom, new.Province, new.City, new.LabelNames, new.Origin)

		new.LongTitle = strings.Trim(new.LongTitle, " ")
		new.LongTitle = strings.Replace(new.LongTitle, " ", "，", -1)
		new.LongTitle = strings.Replace(new.LongTitle, ":", "：", -1)
		new.LongTitle = strings.Replace(new.LongTitle, ",", "，", -1)

		new.Origin = strings.Replace(new.Origin, " ", "", -1)
		new.Origin = strings.Replace(new.Origin, ":", "：", -1)
		new.Origin = strings.Trim(new.Origin, "：")
		new.Origin = strings.Trim(new.Origin, "客户端")

		new.LabelNames = strings.Replace(new.LabelNames, " ", "", -1)
		new.LabelNames = strings.Replace(new.LabelNames, "其他", "", -1)
		new.LabelNames = strings.Replace(new.LabelNames, ",", "，", -1)
		for _, v := range strings.Split(new.LabelNames, "，") {
			if v != "小事件" {
				new.LabelNames = v
				if len([]rune(new.LabelNames)) == 2 {
					break
				}
			}
		}
		new.LabelNames = strings.Split(new.LabelNames, "，")[0]

		if new.Province == "全国" {
			new.Province = "国内"
			new.City = ""
		}

		if new.AreaType == 2 {
			new.Province = "国际"
			new.City = ""
		}

		if strings.Contains(new.Origin, "@") {
			new.Origin = "新浪微博"
		}

		if strings.Contains(new.LabelNames, "明星") || strings.Contains(new.LabelNames, "电影") || strings.Contains(new.LabelNames, "电视剧") || strings.Contains(new.LabelNames, "综艺") || strings.Contains(new.LabelNames, "文化") || strings.Contains(new.LabelNames, "电竞") {
			if new.RatioHotDay < 30 {
				continue
			}
		}

		if len([]rune(new.LongTitle)) <= len([]rune(new.IncidentTitle)) {
			new.IncidentTitle, new.LongTitle = new.LongTitle, new.IncidentTitle

			doc, _ := goquery.NewDocument("http://www.baidu.com/s?rtt=1&bsst=1&cl=2&tn=news&word=" + new.IncidentTitle)

			title, _ := doc.Find("#1 .c-title a").Html()
			title = strings.Replace(title, "\n", "", -1)
			title = strings.Replace(title, "<em>", "", -1)
			title = strings.Replace(title, "</em>", "", -1)
			title = strings.Replace(title, string(regexp.MustCompile("\\s+").Find([]byte(title))), " ", -1)
			title = strings.Trim(title, " ")
			title = strings.Replace(title, " ", "，", -1)
			title = strings.Replace(title, ",", "，", -1)
			title = strings.Replace(title, ":", "：", -1)
			title = strings.Replace(title, "!", "！", -1)
			title = strings.Trim(title, "，")
			title = strings.Trim(title, "：")
			title = strings.Trim(title, "！")
			title = strings.Trim(title, "。")
			title = strings.Replace(title, "?", "？", -1)
			title = strings.Replace(title, "_", "——", -1)
			title = strings.Replace(title, "|", "|", -1)
			title = strings.Replace(title, "...", "……", -1)

			if !strings.Contains(title, "？") && !strings.Contains(title, "——") && !strings.Contains(title, "|") && !strings.Contains(title, "……") {
				new.LongTitle = title
			}
		}

		news = append(news, new)
	}

	return news
}

// GetEvents 获取事件
func GetEvents() []Event {
	events := []Event{}

	data := url.Values{}
	data.Set("sort", "1")
	data.Set("labels", "")
	data.Set("areaType", "")
	data.Set("province", "")
	data.Set("city", "")
	data.Set("startTime", GetDate(-7))
	data.Set("endTime", GetDate(0))
	data.Set("page", "1")
	data.Set("pageSize", "1000")
	data.Set("showTag", "1")
	data.Set("labelShowTag", "1")
	data.Set("webShow", "1")

	fmt.Println(data)

	body, _ := PostForm("http://m.wrd.cn/getBigEventList", data)

	type Resp struct {
		Data []Event `json:"data"`
	}

	var resp Resp
	json.Unmarshal([]byte(body), &resp)

	for i, event := range resp.Data {
		if len(events) == 5 {
			break
		}

		fmt.Println(i, event.Name, event.Province, event.City, event.LabelNames)

		event.LabelNames = strings.Replace(event.LabelNames, " ", "", -1)
		event.LabelNames = strings.Replace(event.LabelNames, "其他", "", -1)
		event.LabelNames = strings.Replace(event.LabelNames, ",", "，", -1)
		for _, v := range strings.Split(event.LabelNames, "，") {
			if v != "小事件" {
				event.LabelNames = v
				if len([]rune(event.LabelNames)) == 2 {
					break
				}
			}
		}
		event.LabelNames = strings.Split(event.LabelNames, "，")[0]

		if event.Province == "全国" {
			event.Province = "国内"
			event.City = ""
		}

		if event.AreaType == 2 {
			event.Province = "国际"
			event.City = ""
		}

		events = append(events, event)
	}

	return events
}

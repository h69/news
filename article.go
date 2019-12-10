package main

import (
	"fmt"
	"strings"
)

// curl -F media=@cover.png "https://api.weixin.qq.com/cgi-bin/material/add_material?type=thumb&access_token="
const cover = "Uo9diUdj5z1RQt3KksOY3w3uEML4hL16n2i446jkUYU"

// Article 文章
type Article struct {
	Title   string
	Digest  string
	Content string
	Cover   string
}

// GenerateArticle 生成文章
func GenerateArticle() Article {
	news := GetNews()
	events := GetEvents()

	var article Article

	article.Title = ""
	article.Digest = "内含 "
	article.Content = ""
	article.Cover = cover

	foreign := 0
	china := make(map[string]int)

	for i, new := range news {
		fmt.Println(i, new.IncidentTitle, new.LongTitle, new.RatioHotDay, new.RatioHotTopCustom, new.Province, new.City, new.LabelNames, new.Origin)

		article.Title += new.IncidentTitle + "；"

		if new.Province == "国际" {
			foreign++
		}

		if new.Province != "国内" && new.Province != "国际" {
			china[new.Province+new.City]++
		}

		article.Content += RenderSubtitle(new.IncidentTitle) + RenderContent(new.Province+new.City+new.LabelNames+"资讯，"+new.LongTitle+"。（"+new.Origin+"）")
	}

	if foreign > 0 {
		article.Digest += IntToString(foreign) + " 封国际电报，"
	}

	article.Digest += IntToString(len(news)-foreign) + " 封国内电报，"

	if len(china) > 0 {
		article.Digest += "其中，"
		for k, v := range china {
			article.Digest += k + " " + IntToString(v) + " 封" + "，"
		}
	}

	article.Digest = strings.Trim(article.Digest, "，")
	article.Digest += "。"

	article.Content = RenderPlaceholder() + RenderQuote(article.Digest) + RenderAuthor("Morse Lab") + RenderCopyright("36Kr") + RenderPlaceholder() + RenderPlaceholder() + article.Content + RenderPlaceholder() + RenderEnd() + RenderPlaceholder()

	more := ""
	for i, event := range events {
		fmt.Println(i, event.Name, event.Province, event.City, event.LabelNames)
		more += RenderCenter(event.Name)
	}

	if len(events) > 0 {
		article.Content += RenderMore() + more + RenderPlaceholder() + RenderPlaceholder() + RenderFooter()
	} else {
		article.Content += RenderPlaceholder() + RenderFooter()
	}

	return article
}

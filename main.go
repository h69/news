package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/robfig/cron"
)

const cover = "Uo9diUdj5z1RQt3KksOY3w3uEML4hL16n2i446jkUYU"

func main() {
	// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	// ./news --id xxx --secret xxx
	flag.StringVar(&AppID, "id", "", "")
	flag.StringVar(&AppSecret, "secret", "", "")
	flag.Parse()

	c := cron.New()
	c.AddFunc("0 0 8 * * ?", func() {
		run()
	})
	c.Start()

	select {}
}

func run() {
	news := GetNews()

	title := ""
	digest := "内含 "
	content := ""

	foreign := 0
	china := make(map[string]int)

	for _, new := range news {
		fmt.Println(new.IncidentTitle, new.LongTitle, new.RatioHotTopCustom, new.RatioHotDay, new.Province, new.City, new.LabelNames, new.Origin)

		title += new.IncidentTitle + "；"

		if new.Province == "国际" {
			foreign++
		}

		if new.Province != "国内" && new.Province != "国际" {
			china[new.Province+new.City]++
		}

		content += PrettifySubtitle(new.IncidentTitle) + PrettifyContent(new.Province+new.City+new.LabelNames+"资讯，"+new.LongTitle+"。（"+new.Origin+"）")
	}

	if foreign > 0 {
		digest += IntToString(foreign) + " 封国际电报，"
	}

	digest += IntToString(len(news)-foreign) + " 封国内电报，"

	if len(china) > 0 {
		digest += "其中，"
		for k, v := range china {
			digest += k + " " + IntToString(v) + " 封" + "，"
		}
	}

	digest = strings.Trim(digest, "，")
	digest += "。"

	content = PrettifyPlaceholder() + PrettifyQuote(digest) + PrettifyPlaceholder() + content + PrettifyPlaceholder() + PrettifyEnd() + PrettifyPlaceholder() + PrettifyCopyright("© 36kr") + PrettifyPlaceholder() + PrettifyFooter()

	// curl -F media=@cover.png "https://api.weixin.qq.com/cgi-bin/material/add_material?type=thumb&access_token="
	AddNews(title, digest, content, cover)
}

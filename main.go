package main

import (
	"flag"

	"github.com/robfig/cron"
)

func main() {
	// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	// ./news --id xxx --secret xxx
	flag.StringVar(&AppID, "id", "", "")
	flag.StringVar(&AppSecret, "secret", "", "")
	flag.Parse()

	c := cron.New()
	c.AddFunc("0 0 8 * * ?", func() {
		article := GenerateArticle()
		AddNews(article.Title, article.Digest, article.Content, article.Cover)
	})
	c.Start()

	select {}
}

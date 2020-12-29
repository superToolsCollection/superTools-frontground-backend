package parser

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"superTools-frontground-backend/global"
	"superTools-frontground-backend/pkg/mq"
)

/**
* @Author: super
* @Date: 2020-09-01 18:59
* @Description:
**/

func ParseArticleUrlList(contents []byte, queueName string, _ string) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		global.Logger.Error(context.Background(), err)
	}

	result := dom.Find("a[rel=bookmark]")
	result.Each(func(i int, selection *goquery.Selection) {
		href, exist := selection.Attr("href")
		if exist {
			global.Logger.Info(context.Background(), "fetching: ", href)
			//将解析到的图书详细信息URL放到消息队列
			//不加延迟会出现问题
			err := mq.Publish(queueName, []byte(href))
			if err != nil {
				global.Logger.Error(context.Background(), err)
			}
		}
	})
}

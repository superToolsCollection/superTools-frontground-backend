package parser

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/crawler/crawerConfig"
	"superTools-frontground-backend/pkg/mq"
	"sync"
)

/**
* @Author: super
* @Date: 2020-08-14 20:49
* @Description:
**/
func ParseTagList(contents []byte, queueName string, url string) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		global.Logger.Error(context.Background(), err)
	}

	result := dom.Find("table[class=tagCol]").Find("a")
	href := ""
	var wg sync.WaitGroup
	result.Each(func(i int, selection *goquery.Selection) {
		href = url + selection.Text()
		for i := 0; i <= 1000; i = i + 20 {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				url := href + "?start=" + strconv.Itoa(i) + "&type=T"
				global.Logger.Info(context.Background(), "fetching: ", url)

				//将解析到的图书详细信息URL放到消息队列
				//不加延迟会出现问题
				err := mq.Publish(queueName, []byte(href))
				if err != nil {
					global.Logger.Error(context.Background(), err)
				}
			}(i)
			//time.Sleep(time.Millisecond * 100)
		}
	})
	wg.Wait()

	err = mq.Publish(queueName, []byte(crawerConfig.StopTAG))
	if err != nil {
		global.Logger.Error(context.Background(), err)
	}
}

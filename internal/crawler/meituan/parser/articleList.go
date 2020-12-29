package parser

import (
	"context"
	"strconv"
	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/crawler/crawerConfig"
	"superTools-frontground-backend/pkg/mq"
)

/**
* @Author: super
* @Date: 2020-09-01 16:00
* @Description:
**/

func ParseArticleList(contents []byte, queueName string, url string) {
	//初始化消息队列
	err := mq.Publish(queueName, []byte(url))
	if err != nil {
		global.Logger.Error(context.Background(), err)
	}
	global.Logger.Info(context.Background(), "fetching: ", url)

	for i := 2; i < 22; i++ {
		url := "https://tech.meituan.com//page/" + strconv.Itoa(i) + ".html"
		global.Logger.Info(context.Background(), "fetching: ", url)
		err := mq.Publish(queueName, []byte(url))
		if err != nil {
			global.Logger.Error(context.Background(), err)
		}
	}

	err = mq.Publish(queueName, []byte(crawerConfig.StopTAG))
	if err != nil {
		global.Logger.Error(context.Background(), err)
	}
}

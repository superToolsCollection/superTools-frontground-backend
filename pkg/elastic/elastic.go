package elastic

import (
	"github.com/olivere/elastic/v7"
	"superTools-frontground-backend/pkg/setting"

	"time"
)

/**
* @Author: super
* @Date: 2020-12-29 11:51
* @Description:
**/


func NewElasticRngine(elasticSetting *setting.ElasticSettingS) (*elastic.Client, error) {
	var client *elastic.Client
	for i := 0; i < 10; i++ {
		// Ping the server by sending a GET request to `/health`.
		var err error
		client, err = elastic.NewClient(
			elastic.SetURL(elasticSetting.Url),
			elastic.SetSniff(false))
		if err == nil{
			return nil, err
		}
		time.Sleep(time.Second)
	}
	return client, nil
}
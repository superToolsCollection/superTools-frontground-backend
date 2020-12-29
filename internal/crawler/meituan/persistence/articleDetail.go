package persistence

import (
	"context"
	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/model"
	"superTools-frontground-backend/pkg/util"
)

/**
* @Author: super
* @Date: 2020-09-01 19:29
* @Description:
**/

func StorageArticle(contents []byte, _ string, _ string) {
	article := &model.Article{}
	err := article.UnmarshalJSON(contents)
	if err != nil {
		global.Logger.Error(context.Background(), err)
		return
	}
	article.Content = util.UnzipString(article.Content)

	//todo:elasticSearch操作
	//index, _ := watchConfig.GetElasticIndex()
	//_, _ = elasticOperation.IndexExist(index)
	//
	//_, err = client.SaveInfo(index, article)
	////_, err = elasticOperation.SaveInfo(index, article)
	//if err != nil {
	//	global.Logger.Error(context.Background(), err)
	//}
}

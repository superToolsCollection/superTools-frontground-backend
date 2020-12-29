package persistence

import (
	"context"
	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/model"
)

/**
* @Author: super
* @Date: 2020-08-16 08:56
* @Description:
**/

func ParseAndStorage(contents []byte, _ string, _ string) {
	book := &model.Book{}
	err := book.UnmarshalJSON(contents)
	if err != nil {
		global.Logger.Error(context.Background(), err)
		return
	}

	//todo：创建dao层进行操作
	//err = DBOperation.InsertBook(book)
	//if err != nil{
	//	global.Logger.Error(context.Background(), err)
	//}
}

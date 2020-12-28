package dao

import (
	"fmt"
	"strings"
	"superTools-frontground-backend/global"
	"superTools-frontground-backend/pkg/db"
	"superTools-frontground-backend/pkg/setting"
	"testing"
)

/**
* @Author: super
* @Date: 2020-12-28 08:59
* @Description:
**/

func TestTagManager_SelectByIDs(t *testing.T) {
	newSetting, err := setting.NewSetting(strings.Split("/Users/super/develop/superTools-frontground-backend/configs", ",")...)
	if err != nil {
		t.Error(err)
	}
	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		t.Error(err)
	}
	conn, err := db.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		t.Error(err)
	}
	tagManager := NewTagManager("tags", conn)

	ids := []string{"1", "2", "3", "4"}
	result, err := tagManager.SelectByIDs(ids)
	if err != nil{
		t.Log(err)
	}
	fmt.Println(result)
}
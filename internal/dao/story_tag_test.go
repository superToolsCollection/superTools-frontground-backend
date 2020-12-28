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
* @Date: 2020-12-28 15:23
* @Description:
**/

func TestStoryTagManager_SelectTagIDsByStoryID(t *testing.T) {
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
	manager := NewStoryTagManager("story_tag", conn)
	result, err := manager.SelectTagIDsByStoryID("464")
	if err != nil {
		t.Log(err)
	}
	fmt.Println(result)
}

func TestStoryTagManager_SelectStoryIDsByTagID(t *testing.T) {
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
	manager := NewStoryTagManager("story_tag", conn)
	result, err := manager.SelectStoryIDsByTagID("3")
	if err != nil {
		t.Log(err)
	}
	fmt.Println(result)
}

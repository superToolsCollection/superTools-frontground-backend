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
* @Date: 2020-12-27 20:09
* @Description:
**/

func TestStoryManager_SelectStory(t *testing.T) {
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
	manager := NewStoryManager("stories", conn)
	story, err := manager.SelectStory()
	if err != nil {
		t.Log(err)
	}
	fmt.Println(story)
}

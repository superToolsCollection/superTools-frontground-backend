package dao

import (
	"strings"
	"superTools-frontground-backend/global"
	"superTools-frontground-backend/pkg/db"
	"superTools-frontground-backend/pkg/setting"
	"testing"
)

/**
* @Author: super
* @Date: 2020-11-18 15:47
* @Description:
**/

func TestProductManager_Insert(t *testing.T) {
	newSetting, err := setting.NewSetting(strings.Split("/Users/super/develop/superTools-backend/configs", ",")...)
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
	productManager := NewProductManager("products", conn)
	product := &Product{
		ID:          2,
		ProductName: "test",
		ProductNum:  12,
		ProductUrl:  "111",
	}
	result, err := productManager.Insert(product)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestProductManager_SelectByKey(t *testing.T) {
	newSetting, err := setting.NewSetting(strings.Split("/Users/super/develop/superTools-backend/configs", ",")...)
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
	productManager := NewProductManager("products", conn)

	product, err := productManager.SelectByKey(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(product)
}

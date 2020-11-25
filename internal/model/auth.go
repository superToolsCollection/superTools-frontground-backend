package model

import (
	"github.com/jinzhu/gorm"
)

/**
* @Author: super
* @Date: 2020-08-22 09:18
* @Description:
**/

type Auth struct {
	*Model
	AppKey    string `gorm:"column:app_key" json:"app_key"`
	AppSecret string `gorm:"column:app_secret" json:"app_secret"`
}

func (a Auth) TableName() string {
	return "auth"
}

//获取表内数据
func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}

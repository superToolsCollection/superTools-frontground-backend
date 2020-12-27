package model

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

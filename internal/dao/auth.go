package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"superTools-frontground-backend/internal/model"
)

/**
* @Author: super
* @Date: 2020-09-23 20:09
* @Description: 在auth表内获取appKey以及appSecret
**/

type Auth struct {
	CreatedOn  string `json:"created_on"`
	CreatedBy  string `json:"created_by"`
	DeletedOn  string `json:"deleted_on"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn string `json:"modified_on"`
	ID         string `json:"id"`
	IsDel      int    `json:"is_del"`
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

type IAuth interface {
	GetAuth(appKey, appSecret string) (*model.Auth, error)
}

type AuthManager struct {
	table string
	conn *gorm.DB
}

func NewAuthManager(table string, conn *gorm.DB) IAuth {
	return &AuthManager{table:table, conn:conn}
}

func (a *AuthManager) GetAuth(appKey, appSecret string) (*model.Auth, error) {
	auth := &model.Auth{}
	result := a.conn.Where("app_key=? and app_secret=?", appKey, appSecret).Find(auth)
	if result.RecordNotFound(){
		return nil, errors.New("wrong appKey or appSecret")
	}
	return auth, nil
}

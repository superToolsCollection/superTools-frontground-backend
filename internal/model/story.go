package model

import (
	"superTools-frontground-backend/pkg/app"
)

/**
* @Author: super
* @Date: 2020-09-16 07:38
* @Description: 与stories对应的结构体
* @Group: BedTimeStory
**/
type Story struct {
	*Model
	Author string `gorm:"column:author" json:"author"`
	State  uint8  `gorm:"column:state" json:"state"`
	Story  string `gorm:"column:story" json:"story"`
}

// TableName sets the insert table name for this struct type
func (b Story) TableName() string {
	return "stories"
}

//用于swagger的内容展示
type BedtimeStorySwagger struct {
	List  []*Story
	Pager *app.Pager
}

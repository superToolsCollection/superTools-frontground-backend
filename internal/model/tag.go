package model

import (
	"superTools-frontground-backend/pkg/app"
)

/**
* @Author: super
* @Date: 2020-09-16 07:40
* @Description: 与tags表对应的结构体
* @Group: BedTimeStory
**/

type Tag struct {
	*Model
	Name  string `gorm:"column:name" json:"name"`
	State uint8  `gorm:"column:state" json:"state"`
}

// TableName sets the insert table name for this struct type
func (s Tag) TableName() string {
	return "tags"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

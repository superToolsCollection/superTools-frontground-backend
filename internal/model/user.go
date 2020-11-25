package model

/**
* @Author: super
* @Date: 2020-11-21 11:24
* @Description:
**/

type User struct {
	ID           string `gorm:"column:id;primary_key" json:"id"`
	NickName     string `gorm:"column:nick_name" json:"nick_name"`
	UserName     string `gorm:"column:user_name" json:"user_name"`
	HashPassword string `gorm:"column:hash_password" json:"-"`
}

func (u User) TableName() string {
	return "users"
}

package dao

import (
	"errors"

	"superTools-frontground-backend/internal/model"

	"github.com/jinzhu/gorm"
)

/**
* @Author: super
* @Date: 2020-11-21 11:25
* @Description:
**/

type User struct {
	ID           string `json:"id"`
	NickName     string `json:"nick_name"`
	UserName     string `json:"user_name"`
	HashPassword string `json:"-"`
}

type IUser interface {
	Insert(user *User) (userID string, err error)
	Delete(id string) bool
	Update(user *User) error
	SelectByUserName(userName string) (result *model.User, err error)
	SelectByID(id string) (result *model.User, err error)
	SelectAll() ([]*model.User, error)
}

type UserManager struct {
	table string
	conn  *gorm.DB
}

func NewUserManager(table string, conn *gorm.DB) IUser {
	return &UserManager{table: table, conn: conn}
}

func (m *UserManager) Insert(user *User) (userID string, err error) {
	u := &model.User{
		ID:           user.ID,
		UserName:     user.UserName,
		NickName:     user.NickName,
		HashPassword: user.HashPassword,
	}
	result := m.conn.Create(u)
	if result.RowsAffected == int64(0) {
		return "", errors.New("insert error")
	}
	return u.ID, nil
}

func (m *UserManager) Delete(id string) bool {
	result := m.conn.Where("id=?", id).Delete(model.User{})
	if result.RowsAffected == int64(0) {
		return false
	}
	return true
}

func (m *UserManager) Update(user *User) error {
	u := &model.User{
		ID:           user.ID,
		UserName:     user.UserName,
		NickName:     user.NickName,
		HashPassword: user.HashPassword,
	}
	result := m.conn.Model(u).Where("id=?", u.ID).Updates(u)
	if result.RowsAffected == int64(0) {
		return errors.New("update error")
	}
	return nil
}

func (m *UserManager) SelectByUserName(userName string) (result *model.User, err error) {
	result = &model.User{}
	r := m.conn.Where("user_name = ?", userName).Find(result)
	if r.RecordNotFound() {
		return nil, errors.New("wrong userName")
	}
	return result, nil
}

func (m *UserManager) SelectAll() ([]*model.User, error) {
	var users []*model.User
	if err := m.conn.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (m *UserManager) SelectByID(id string) (result *model.User, err error) {
	result = &model.User{}
	r := m.conn.Where("id = ?", id).Find(result)
	if r.RecordNotFound() {
		return nil, errors.New("wrong id")
	}
	return result, nil
}

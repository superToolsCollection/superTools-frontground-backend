package dao

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"superTools-frontground-backend/internal/model"
)

/**
* @Author: super
* @Date: 2020-09-22 09:36
* @Description: 用于操作story_tag表
**/
type Tag struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type ITag interface {
	SelectByID(id string) (Tag, error)
	SelectByIDs(id []string) ([]Tag, error)
}

type TagManager struct {
	table string
	conn  *gorm.DB
}

func NewTagManager(table string, conn *gorm.DB) ITag {
	return &TagManager{table: table, conn: conn}
}

func (m *TagManager) SelectByID(id string) (Tag, error) {
	t := &model.Tag{}
	result := m.conn.Where("id=?", id).Find(t)
	if result.RecordNotFound() {
		return Tag{}, errors.New("select tag by id error")
	}
	return Tag{
		ID:    t.ID,
		Name:  t.Name,
		State: t.State,
	}, nil
}

func (m *TagManager) SelectByIDs(ids []string) ([]Tag, error) {
	var tags []*model.Tag
	result := m.conn.Where("id in (?)", ids).Find(&tags)
	if result.RecordNotFound() {
		return nil, errors.New("select tags by id error")
	}
	t := make([]Tag, 0)
	for _, tag := range tags {
		temp := Tag{
			ID:    tag.ID,
			Name:  tag.Name,
			State: tag.State,
		}
		t = append(t, temp)
	}

	return t, nil
}

func (t Tag) String() string {
	return fmt.Sprintf("id: %s, name: %s, state:%d\n", t.ID, t.Name, t.State)
}

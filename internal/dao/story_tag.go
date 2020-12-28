package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"superTools-frontground-backend/internal/model"
)

/**
* @Author: super
* @Date: 2020-09-22 14:14
* @Description: 用于操作story_tag_map表
**/
type IStoryTag interface {
	SelectTagIDsByStoryID(id string) ([]string, error)
	SelectStoryIDsByTagID(id string) ([]string, error)
}

type StoryTagManager struct {
	table string
	conn  *gorm.DB
}

func NewStoryTagManager(table string, conn *gorm.DB) IStoryTag {
	return &StoryTagManager{table: table, conn: conn}
}

func (m *StoryTagManager) SelectTagIDsByStoryID(id string) ([]string, error) {
	var st []*model.StoryTagMap
	result := m.conn.Where("story_id=?", id).Find(&st)
	err := result.Error
	if err != nil {
		return nil, errors.New("SelectTagIDsByStoryID error")
	} else if result.RecordNotFound() {
		return []string{}, nil
	}
	ids := []string{}
	for _, stm := range st {
		ids = append(ids, stm.TagID)
	}
	return ids, nil
}

func (m *StoryTagManager) SelectStoryIDsByTagID(id string) ([]string, error) {
	var st []*model.StoryTagMap
	result := m.conn.Where("tag_id=?", id).Find(&st)
	err := result.Error
	if err != nil {
		return nil, errors.New("SelectStoryIDsByTagID error")
	} else if result.RecordNotFound() {
		return []string{}, nil
	}
	ids := []string{}
	for _, stm := range st {
		ids = append(ids, stm.StoryID)
	}
	return ids, nil
}

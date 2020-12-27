package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"superTools-frontground-backend/internal/model"
)

/**
* @Author: super
* @Date: 2020-09-22 14:06
* @Description: 用于操作stories表
**/

type Story struct {
	ID     string `json:"id"`
	TagID  string `json:"tag_id"`
	Story  string `json:"story"`
	Author string `json:"author"`
	State  uint8  `json:"state"`
}

type IStory interface {
	SelectStory() (*Story, error)
}

type StoryManager struct {
	table string
	conn  *gorm.DB
}

func NewStoryManager(table string, conn *gorm.DB) IStory {
	return &StoryManager{table: table, conn: conn}
}

//从数据库中随机选择一条数据返回
func (m *StoryManager) SelectStory() (*Story, error) {
	story := &model.Story{}
	result := m.conn.Where().Find(story)
	if result.RecordNotFound() {
		return nil, errors.New("get story error")
	}
	return &Story{
		ID:     story.ID,
		Story:  story.Story,
		Author: story.Author,
		State:  story.State,
	}, nil
}

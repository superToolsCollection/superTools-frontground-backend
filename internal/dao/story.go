package dao

import (
	"github.com/jinzhu/gorm"
)

/**
* @Author: super
* @Date: 2020-09-22 14:06
* @Description: 用于操作stories表
**/

type Story struct {
	ID     string `json:"id"`
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
	result := m.conn.Raw(`SELECT s1.id AS id, s1.author AS author, s1.story AS story, s1.state as state FROM stories AS s1 JOIN (SELECT ROUND(RAND() * ((SELECT MAX(id) FROM stories) - (SELECT MIN(id) FROM stories)) + (SELECT MIN(id) FROM stories)) AS id) AS s2 ON s1.id >= s2.id LIMIT 1;`).Row()
	story := &Story{}
	err := result.Scan(&story.ID, &story.Author, &story.Story, &story.State)
	if err != nil {
		return nil, err
	}
	return story, nil
}

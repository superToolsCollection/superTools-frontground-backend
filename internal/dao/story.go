package dao

import (
	"superTools-frontground-backend/internal/model"
	"superTools-frontground-backend/pkg/app"
)

/**
* @Author: super
* @Date: 2020-09-22 14:06
* @Description: 用于操作stories表
**/

type Story struct {
	ID         uint32 `json:"id"`
	TagID      uint32 `json:"tag_id"`
	Story      string `json:"story"`
	Author     string `json:"author"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      uint8  `json:"state"`
}

func (d *Dao) CreateStory(param *Story) (*model.BedtimeStory, error) {
	story := model.BedtimeStory{
		Story:  param.Story,
		Author: param.Author,
		State:  param.State,
		Model:  &model.Model{CreatedBy: param.CreatedBy},
	}
	return story.Create(d.engine)
}

func (d *Dao) UpdateStory(param *Story) error {
	story := model.BedtimeStory{Model: &model.Model{ID: param.ID}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}

	if param.Story != "" {
		values["story"] = param.Story
	}

	return story.Update(d.engine, values)
}

func (d *Dao) GetStory(id uint32, state uint8) (model.BedtimeStory, error) {
	story := model.BedtimeStory{Model: &model.Model{ID: id}, State: state}
	return story.Get(d.engine)
}

func (d *Dao) GetStoryOnly(id uint32, state uint8) (model.BedtimeStory, error) {
	story := model.BedtimeStory{Model: &model.Model{ID: id}, State: state}
	return story.Get(d.engine)
}

func (d *Dao) DeleteStory(id uint32) error {
	story := model.BedtimeStory{Model: &model.Model{ID: id}}
	return story.Delete(d.engine)
}

func (d *Dao) CountStoryListByTagID(id uint32, state uint8) (int, error) {
	story := model.BedtimeStory{State: state}
	return story.CountByTagID(d.engine, id)
}

func (d *Dao) GetStoryListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.StoryRow, error) {
	story := model.BedtimeStory{State: state}
	return story.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
}

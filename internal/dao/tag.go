package dao

import (
	"superTools-frontground-backend/internal/model"
	"superTools-frontground-backend/pkg/app"
)

/**
* @Author: super
* @Date: 2020-09-22 09:36
* @Description: 用于操作story_tag表
**/

func (d *Dao) GetTag(id string, state uint8) (model.StoryTag, error) {
	tag := model.StoryTag{Model: &model.Model{ID: id}, State: state}
	return tag.Get(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.StoryTag, error) {
	tag := model.StoryTag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetTagListByIDs(ids []uint32, state uint8) ([]*model.StoryTag, error) {
	tag := model.StoryTag{State: state}
	return tag.ListByIDs(d.engine, ids)
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.StoryTag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.StoryTag{
		Name:  name,
		State: state,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id string, name string, state uint8, modifiedBy string) error {
	tag := model.StoryTag{
		Model: &model.Model{
			ID: id,
		},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id string) error {
	tag := model.StoryTag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}

package dao

import (
	"superTools-frontground-backend/internal/model"
)

/**
* @Author: super
* @Date: 2020-09-22 14:14
* @Description: 用于操作story_tag_map表
**/

func (d *Dao) GetStoryTagByAID(storyID uint32) (model.StoryTagMap, error) {
	articleTag := model.StoryTagMap{StoryID: storyID}
	return articleTag.GetByAID(d.engine)
}

func (d *Dao) GetStoryTagListByTID(tagID uint32) ([]*model.StoryTagMap, error) {
	articleTag := model.StoryTagMap{TagID: tagID}
	return articleTag.ListByTID(d.engine)
}

func (d *Dao) GetStoryTagListByAIDs(articleIDs []uint32) ([]*model.StoryTagMap, error) {
	articleTag := model.StoryTagMap{}
	return articleTag.ListByAIDs(d.engine, articleIDs)
}

func (d *Dao) CreateStoryTag(storyID, tagID uint32, createdBy string) error {
	articleTag := model.StoryTagMap{
		Model: &model.Model{
			CreatedBy: createdBy,
		},
		StoryID: storyID,
		TagID:   tagID,
	}
	return articleTag.Create(d.engine)
}

func (d *Dao) UpdateStoryTag(storyID, tagID uint32, modifiedBy string) error {
	articleTag := model.StoryTagMap{StoryID: storyID}
	values := map[string]interface{}{
		"story_id":    storyID,
		"tag_id":      tagID,
		"modified_by": modifiedBy,
	}
	return articleTag.UpdateOne(d.engine, values)
}

func (d *Dao) DeleteStoryTag(storyID uint32) error {
	articleTag := model.StoryTagMap{StoryID: storyID}
	return articleTag.DeleteOne(d.engine)
}

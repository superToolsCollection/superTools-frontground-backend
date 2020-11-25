package model

import (
	"github.com/jinzhu/gorm"
)

/**
* @Author: super
* @Date: 2020-09-16 07:41
* @Description: 与story_tag_map对应的结构体
* @Group: BedTimeStory
**/

type StoryTagMap struct {
	*Model
	StoryID uint32 `gorm:"column:story_id" json:"story_id"`
	TagID   uint32 `gorm:"column:tag_id" json:"tag_id"`
}

// TableName sets the insert table name for this struct type
func (s StoryTagMap) TableName() string {
	return "story_tag_map"
}

//以下内容为CRUD操作
func (s StoryTagMap) GetByAID(db *gorm.DB) (StoryTagMap, error) {
	var articleTag StoryTagMap
	err := db.Where("story_id = ? AND is_del = ?", s.StoryID, 0).First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articleTag, err
	}

	return articleTag, nil
}

func (s StoryTagMap) ListByTID(db *gorm.DB) ([]*StoryTagMap, error) {
	var articleTags []*StoryTagMap
	if err := db.Where("tag_id = ? AND is_del = ?", s.TagID, 0).Find(&articleTags).Error; err != nil {
		return nil, err
	}

	return articleTags, nil
}

func (s StoryTagMap) ListByAIDs(db *gorm.DB, articleIDs []uint32) ([]*StoryTagMap, error) {
	var articleTags []*StoryTagMap
	err := db.Where("story_id IN (?) AND is_del = ?", articleIDs, 0).Find(&articleTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articleTags, nil
}

func (s StoryTagMap) Create(db *gorm.DB) error {
	if err := db.Create(&s).Error; err != nil {
		return err
	}

	return nil
}

func (s StoryTagMap) UpdateOne(db *gorm.DB, values interface{}) error {
	if err := db.Model(&s).Where("story_id = ? AND is_del = ?", s.StoryID, 0).Limit(1).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (s StoryTagMap) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", s.Model.ID, 0).Delete(&s).Error; err != nil {
		return err
	}

	return nil
}

func (s StoryTagMap) DeleteOne(db *gorm.DB) error {
	if err := db.Where("story_id = ? AND is_del = ?", s.StoryID, 0).Delete(&s).Limit(1).Error; err != nil {
		return err
	}

	return nil
}

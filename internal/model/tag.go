package model

import (
	"github.com/jinzhu/gorm"

	"superTools-frontground-backend/pkg/app"
)

/**
* @Author: super
* @Date: 2020-09-16 07:40
* @Description: 与story_tag表对应的结构体
* @Group: BedTimeStory
**/

type StoryTag struct {
	*Model
	Name  string `gorm:"column:name" json:"name"`
	State uint8  `gorm:"column:state" json:"state"`
}

// TableName sets the insert table name for this struct type
func (s StoryTag) TableName() string {
	return "story_tag"
}

type TagSwagger struct {
	List  []*StoryTag
	Pager *app.Pager
}

//以下是数据库的CRUD操作
func (s StoryTag) Count(db *gorm.DB) (int, error) {
	var count int
	if s.Name != "" {
		db = db.Where("name = ?", s.Name)
	}
	db = db.Where("state = ?", s.State)
	if err := db.Model(&s).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (s StoryTag) List(db *gorm.DB, pageOffset, pageSize int) ([]*StoryTag, error) {
	var tags []*StoryTag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if s.Name != "" {
		db = db.Where("name = ?", s.Name)
	}
	db = db.Where("state = ?", s.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (s StoryTag) ListByIDs(db *gorm.DB, ids []uint32) ([]*StoryTag, error) {
	var tags []*StoryTag
	db = db.Where("state = ? AND is_del = ?", s.State, 0)
	err := db.Where("id IN (?)", ids).Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

func (s StoryTag) Get(db *gorm.DB) (StoryTag, error) {
	var tag StoryTag
	err := db.Where("id = ? AND is_del = ? AND state = ?", s.ID, 0, s.State).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}

	return tag, nil
}

func (s StoryTag) Create(db *gorm.DB) error {
	return db.Create(&s).Error
}

func (s StoryTag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&s).Where("id = ? AND is_del = ?", s.ID, 0).Updates(values).Error
}

func (s StoryTag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", s.Model.ID, 0).Delete(&s).Error
}

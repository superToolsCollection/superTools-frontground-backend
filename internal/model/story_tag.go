package model

/**
* @Author: super
* @Date: 2020-12-28 15:29
* @Description:
**/

type StoryTagMap struct {
	*Model
	StoryID string `gorm:"column:story_id" json:"story_id"`
	TagID   string `gorm:"column:tag_id" json:"tag_id"`
}

// TableName sets the insert table name for this struct type
func (s StoryTagMap) TableName() string {
	return "story_tag_map"
}

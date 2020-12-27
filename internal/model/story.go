package model

import (
	"github.com/jinzhu/gorm"

	"superTools-frontground-backend/pkg/app"
)

/**
* @Author: super
* @Date: 2020-09-16 07:38
* @Description: 与stories对应的结构体
* @Group: BedTimeStory
**/
type Story struct {
	*Model
	Author string `gorm:"column:author" json:"author"`
	State  uint8  `gorm:"column:state" json:"state"`
	Story  string `gorm:"column:story" json:"story"`
}

// TableName sets the insert table name for this struct type
func (b Story) TableName() string {
	return "stories"
}

//用于swagger的内容展示
type BedtimeStorySwagger struct {
	List  []*Story
	Pager *app.Pager
}

//以下内容是数据库的CRUD操作
func (b Story) Create(db *gorm.DB) (*Story, error) {
	if err := db.Create(&b).Error; err != nil {
		return nil, err
	}

	return &b, nil
}

func (b Story) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&b).Where("id = ? AND is_del = ?", b.ID, 0).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (b Story) Get(db *gorm.DB) (Story, error) {
	var story Story
	db = db.Where("id = ? AND state = ? AND is_del = ?", b.ID, b.State, 0)
	err := db.First(&story).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return story, err
	}

	return story, nil
}

func (b Story) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", b.Model.ID, 0).Delete(&b).Error; err != nil {
		return err
	}

	return nil
}

type StoryRow struct {
	StoryID string
	TagID   string
	TagName string
	Story   string
}

func (b Story) ListByTagID(db *gorm.DB, tagID string, pageOffset, pageSize int) ([]*StoryRow, error) {
	fields := []string{"st.id AS story_id", "st.story"}
	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Table(StoryTagMap{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+StoryTag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Story{}.TableName()+"` AS st ON at.story_id = st.id").
		Where("at.`tag_id` = ? AND st.state = ? AND st.is_del = ?", tagID, b.State, 0).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*StoryRow
	for rows.Next() {
		r := &StoryRow{}
		if err := rows.Scan(&r.StoryID, &r.Story, &r.TagID, &r.TagName); err != nil {
			return nil, err
		}

		articles = append(articles, r)
	}

	return articles, nil
}

func (b Story) CountByTagID(db *gorm.DB, tagID string) (int, error) {
	var count int
	err := db.Table(StoryTag{}.TableName()+" AS st").
		Joins("LEFT JOIN `"+StoryTagMap{}.TableName()+"` AS stm ON stm.tag_id = st.id").
		Joins("LEFT JOIN `"+Story{}.TableName()+"` AS bs ON stm.story_id = bs.id").
		Where("stm.`tag_id` = ? AND st.state = ? AND st.is_del = ?", tagID, b.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

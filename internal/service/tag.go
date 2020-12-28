package service

import (
	"strings"
	"superTools-frontground-backend/internal/dao"
)

/**
* @Author: super
* @Date: 2020-09-18 15:00
* @Description: tag相关内容入参验证与service代码
**/
type SelectTagRequest struct {
	ID string `form:"id" binding:"required"`
}

type SelectTagsRequest struct {
	IDs string `form:"ids" binding:"required"`
}

type Tag struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type ITagService interface {
	GetTag(param *SelectTagRequest) (*Tag, error)
	GetTags(param *SelectTagsRequest) ([]*Tag, error)
}

type TagService struct {
	tagDao dao.ITag
}

func NewTagServcice(tagDao dao.ITag) ITagService {
	return &TagService{tagDao: tagDao}
}

func (s *TagService) GetTag(param *SelectTagRequest) (*Tag, error) {
	result, err := s.tagDao.SelectByID(param.ID)
	if err != nil {
		return nil, err
	}
	return &Tag{
		ID:    result.ID,
		Name:  result.Name,
		State: result.State,
	}, nil
}

func (s *TagService) GetTags(param *SelectTagsRequest) ([]*Tag, error) {
	//todo:确认传递参数
	ids := strings.Split(strings.TrimSpace(param.IDs), ",")
	result, err := s.tagDao.SelectByIDs(ids)
	if err != nil {
		return nil, err
	}
	tags := make([]*Tag, len(result))
	for i, r := range result {
		tags[i] = &Tag{
			ID:    r.ID,
			Name:  r.Name,
			State: r.State,
		}
	}
	return tags, nil
}

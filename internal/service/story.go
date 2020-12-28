package service

import "superTools-frontground-backend/internal/dao"

/**
* @Author: super
* @Date: 2020-09-18 15:05
* @Description: story相关内容入参验证与service代码
**/

type Story struct {
	ID     string `json:"id"`
	Story  string `json:"story"`
	Author string `json:"author"`
	State  uint8  `json:"state"`
	Tags   []*Tag `json:"tags"`
}

type IStoryService interface {
	GetStory() (*Story, error)
}

type StoryService struct {
	storyDao    dao.IStory
	tagDao      dao.ITag
	storyTagDao dao.IStoryTag
}

func NewStoryService(storyDao dao.IStory, tagDao dao.ITag, storyTagDao dao.IStoryTag) IStoryService {
	return &StoryService{storyDao: storyDao, tagDao: tagDao, storyTagDao: storyTagDao}
}

func (s *StoryService) GetStory() (*Story, error) {
	result, err := s.storyDao.SelectStory()
	if err != nil {
		return nil, err
	}
	tagIds, err := s.storyTagDao.SelectTagIDsByStoryID(result.ID)
	if err != nil {
		return nil, err
	}
	daoTags, err := s.tagDao.SelectByIDs(tagIds)
	if err != nil {
		return nil, err
	}
	tags := make([]*Tag, len(daoTags))
	for i, t := range daoTags {
		tags[i] = &Tag{
			ID:    t.ID,
			Name:  t.Name,
			State: t.State,
		}
	}
	return &Story{
		ID:     result.ID,
		Story:  result.Story,
		Author: result.Author,
		State:  result.State,
		Tags:   tags,
	}, nil
}

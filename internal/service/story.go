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
	Tag    *Tag   `json:"tag"`
}

type IStoryService interface {
	GetStory() (*Story, error)
}

type StoryService struct {
	storyDao dao.IStory
}

func NewStoryService(storyDao dao.IStory) IStoryService {
	return &StoryService{storyDao: storyDao}
}

func (s *StoryService) GetStory() (*Story, error) {
	result, err := s.storyDao.SelectStory()
	if err != nil {
		return nil, err
	}
	return &Story{
		ID:     result.ID,
		Story:  result.Story,
		Author: result.Author,
		State:  result.State,
	}, nil
}

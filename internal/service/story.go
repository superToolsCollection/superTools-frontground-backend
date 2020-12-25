package service

import (
	"superTools-frontground-backend/internal/dao"
	"superTools-frontground-backend/internal/model"
	"superTools-frontground-backend/pkg/app"
	"superTools-frontground-backend/pkg/util"
)

/**
* @Author: super
* @Date: 2020-09-18 15:05
* @Description: story相关内容入参验证与service代码
**/

type StoryRequest struct {
	ID    string `form:"id" binding:"required,min=2,max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type StoryListRequest struct {
	TagID string `form:"tag_id" binding:"required,min=2,max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateStoryRequest struct {
	TagID     string `form:"tag_id" binding:"required,min=2,max=100"`
	Content   string `form:"content" binding:"required,min=2,max=4294967295"`
	Author    string `form:"author" binding:"required,min=2,max=4294967295"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateStoryRequest struct {
	ID         string `form:"id" binding:"required,min=2,max=100"`
	TagID      string `form:"tag_id" binding:"required,min=2,max=100"`
	Content    string `form:"content" binding:"min=2,max=4294967295"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteStoryRequest struct {
	ID string `form:"id" binding:"required,min=2,max=100"`
}

type Story struct {
	ID     string          `json:"id"`
	Story  string          `json:"story"`
	Author string          `json:"author"`
	State  uint8           `json:"state"`
	Tag    *model.StoryTag `json:"tag"`
}

type StoryOnly struct {
	Story  string `json:"story"`
	Author string `json:"author"`
}

func (svc *Service) GetStory(param *StoryRequest) (*Story, error) {
	story, err := svc.dao.GetStory(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	storyTagMap, err := svc.dao.GetStoryTagByAID(story.ID)
	if err != nil {
		return nil, err
	}

	tag, err := svc.dao.GetTag(storyTagMap.TagID, 1) //todo: state魔法数
	if err != nil {
		return nil, err
	}

	//将base64编码后的字符串返回给前端，提升传输效率
	content, err := util.EncodeBase64(story.Story)
	if err != nil {
		return nil, err
	}
	return &Story{
		ID:     story.ID,
		Story:  content,
		Author: story.Author,
		State:  story.State,
		Tag:    &tag,
	}, nil
}

func (svc *Service) GetStoryOnly(param *StoryRequest) (*StoryOnly, error) {
	story, err := svc.dao.GetStoryOnly(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	//将base64编码后的字符串返回给前端，提升传输效率
	content, err := util.EncodeBase64(story.Story)
	if err != nil {
		return nil, err
	}
	return &StoryOnly{
		Story:  content,
		Author: story.Author,
	}, nil
}

func (svc *Service) GetStoryList(param *StoryListRequest, pager *app.Pager) ([]*Story, int, error) {
	count, err := svc.dao.CountStoryListByTagID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}

	storyRows, err := svc.dao.GetStoryListByTagID(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var stories []*Story
	for _, story := range storyRows {
		//将base64编码后的字符串返回给前端，提升传输效率
		content, err := util.EncodeBase64(story.Story)
		if err != nil {
			return nil, 0, err
		}
		stories = append(stories, &Story{
			ID:    story.StoryID,
			Story: content,
			Tag:   &model.StoryTag{Model: &model.Model{ID: story.TagID}, Name: story.TagName},
		})
	}

	return stories, count, nil
}

func (svc *Service) CreateStory(param *CreateStoryRequest) error {
	//TODO 做内容校验，相同的内容就不再插入

	//前端传递过来是base64编码后的字符串
	content := util.DecodeBase64(param.Content)
	story, err := svc.dao.CreateStory(&dao.Story{
		Story:     content,
		Author:    param.Author,
		State:     param.State,
		CreatedBy: param.CreatedBy,
	})
	if err != nil {
		return err
	}

	err = svc.dao.CreateStoryTag(story.ID, param.TagID, param.CreatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) UpdateStory(param *UpdateStoryRequest) error {
	//TODO 做内容校验，相同的内容就不再插入
	//前端传递过来是base64编码后的字符串
	content := util.DecodeBase64(param.Content)
	err := svc.dao.UpdateStory(&dao.Story{
		ID:         param.ID,
		Story:      content,
		State:      param.State,
		ModifiedBy: param.ModifiedBy,
	})
	if err != nil {
		return err
	}

	err = svc.dao.UpdateStoryTag(param.ID, param.TagID, param.ModifiedBy)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) DeleteStory(param *DeleteStoryRequest) error {
	err := svc.dao.DeleteStory(param.ID)
	if err != nil {
		return err
	}

	err = svc.dao.DeleteStoryTag(param.ID)
	if err != nil {
		return err
	}

	return nil
}

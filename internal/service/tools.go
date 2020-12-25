package service

import (
	"superTools-frontground-backend/internal/dao"
	"superTools-frontground-backend/pkg/app"
)

/**
* @Author: super
* @Date: 2020-11-25 15:06
* @Description:
**/
type GetToolByKeyRequest struct {
	ID string `form:"id" binding:"required,min=2,max=4294967295"`
}

type GetToolByNameRequest struct {
	Name string `form:"name" binding:"required,min=2,max=4294967295"`
}

type Tool struct {
	CreateOn    string `json:"create_on"`
	CreatedBy   string `json:"created_by"`
	ID          string `json:"id"`
	ModifiedBy  string `json:"modified_by"`
	APIDescribe string `json:"api_describe"`
	DeleteOn    string `json:"delete_on"`
	ModifiedOn  string `json:"modified_on"`
	Name        string `json:"name"`
	State       int    `json:"state"`
	API         string `json:"api"`
}

type IToolService interface {
	GetToolByKey(param *GetToolByKeyRequest) (*Tool, error)
	GetToolByName(param *GetToolByNameRequest) (*Tool, error)
	GetToolList(pager *app.Pager) ([]*Tool, error)
	GetAllTools() ([]*Tool, error)
}

type ToolService struct {
	toolDao dao.ITool
}

func NewToolService(toolDao dao.ITool) IToolService {
	return &ToolService{toolDao: toolDao}
}

func (s *ToolService) GetToolByKey(param *GetToolByKeyRequest) (*Tool, error) {
	result, err := s.toolDao.SelectByKey(param.ID)
	if err != nil {
		return nil, err
	}
	tool := &Tool{
		ID:          result.ID,
		Name:        result.Name,
		API:         result.API,
		APIDescribe: result.APIDescribe,
		CreatedBy:   result.CreatedBy,
		CreateOn:    result.CreatedOn,
		ModifiedBy:  result.ModifiedBy,
		ModifiedOn:  result.ModifiedOn,
		DeleteOn:    result.DeletedOn,
		State:       result.State,
	}
	return tool, nil
}

func (s *ToolService) GetToolByName(param *GetToolByNameRequest) (*Tool, error) {
	result, err := s.toolDao.SelectByName(param.Name)
	if err != nil {
		return nil, err
	}
	tool := &Tool{
		ID:          result.ID,
		Name:        result.Name,
		API:         result.API,
		APIDescribe: result.APIDescribe,
		CreatedBy:   result.CreatedBy,
		CreateOn:    result.CreatedOn,
		ModifiedBy:  result.ModifiedBy,
		ModifiedOn:  result.ModifiedOn,
		DeleteOn:    result.DeletedOn,
		State:       result.State,
	}
	return tool, nil
}

func (s *ToolService) GetToolList(pager *app.Pager) ([]*Tool, error) {
	tools, err := s.toolDao.SelectList(pager.Page, pager.PageSize)
	if err != nil {
		return nil, err
	}
	var toolList []*Tool
	for _, tool := range tools {
		toolList = append(toolList, &Tool{
			ID:          tool.ID,
			Name:        tool.Name,
			API:         tool.API,
			APIDescribe: tool.APIDescribe,
			CreatedBy:   tool.CreatedBy,
			CreateOn:    tool.CreatedOn,
			ModifiedBy:  tool.ModifiedBy,
			ModifiedOn:  tool.ModifiedOn,
			DeleteOn:    tool.DeletedOn,
			State:       tool.State,
		})
	}
	return toolList, nil
}

func (s *ToolService) GetAllTools()([]*Tool, error){
	tools, err := s.toolDao.SelectAll()
	if err != nil {
		return nil, err
	}
	var toolList []*Tool
	for _, tool := range tools {
		toolList = append(toolList, &Tool{
			ID:          tool.ID,
			Name:        tool.Name,
			API:         tool.API,
			APIDescribe: tool.APIDescribe,
			CreatedBy:   tool.CreatedBy,
			CreateOn:    tool.CreatedOn,
			ModifiedBy:  tool.ModifiedBy,
			ModifiedOn:  tool.ModifiedOn,
			DeleteOn:    tool.DeletedOn,
			State:       tool.State,
		})
	}
	return toolList, nil
}
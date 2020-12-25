package dao

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"superTools-frontground-backend/internal/model"
	"superTools-frontground-backend/pkg/app"
)

/**
* @Author: super
* @Date: 2020-12-25 20:50
* @Description:
**/

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

type ITool interface {
	Insert(tool *Tool) (string, error)
	Delete(id string) error
	Update(tool *Tool) error
	SelectByKey(id string) (*model.Tool, error)
	SelectAll() ([]*model.Tool, error)
	SelectByName(name string) (*model.Tool, error)
	SelectList(page, pageSize int) ([]*model.Tool, error)
}

type ToolManager struct {
	table string
	conn  *gorm.DB
}

func (m *ToolManager) Insert(tool *Tool) (string, error) {
	t := &model.Tool{
		Model: &model.Model{
			ID:        tool.ID,
			CreatedBy: tool.CreatedBy,
		},
		APIDescribe: tool.APIDescribe,
		Name:        tool.Name,
		State:       tool.State,
		API:         tool.API,
	}
	result := m.conn.Create(t)
	if result.RowsAffected == int64(0) {
		return "", errors.New(fmt.Sprintf("dao.insertTool err: %v", result.Error.Error()))
	}
	return t.ID, nil
}

func (m *ToolManager) Delete(id string) error {
	result := m.conn.Where("id=?", id).Delete(model.Tool{})
	if result.RowsAffected == int64(0) {
		return errors.New(fmt.Sprintf("delete error: %v", result.Error.Error()))
	}
	return nil
}

func (m *ToolManager) Update(tool *Tool) error {
	t := &model.Tool{
		Model: &model.Model{
			ID:         tool.ID,
			ModifiedBy: tool.ModifiedBy,
		},
		APIDescribe: tool.APIDescribe,
		Name:        tool.Name,
		State:       tool.State,
		API:         tool.API,
	}
	fmt.Println("update", t.Name, t.State, t.ID)
	result := m.conn.Model(t).Where("id=?", t.ID).Updates(t)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("update error: %v", result.Error.Error()))
	}
	return nil
}

func (m *ToolManager) SelectByKey(id string) (*model.Tool, error) {
	t := &model.Tool{}
	result := m.conn.Where("id=?", id).Find(t)
	if result.RecordNotFound() {
		return nil, errors.New(fmt.Sprintf("SelectByKey error: %v", result.Error.Error()))
	}
	return t, nil
}

func (m *ToolManager) SelectAll() ([]*model.Tool, error) {
	var tools []*model.Tool
	if err := m.conn.Find(&tools).Error; err != nil {
		return nil, err
	}
	return tools, nil
}

func (m *ToolManager) SelectByName(name string) (*model.Tool, error) {
	t := &model.Tool{}
	result := m.conn.Where("name=?", name).Find(t)
	if result.RecordNotFound() {
		return nil, errors.New(fmt.Sprintf("SelectByName error: %v", result.Error.Error()))
	}
	return t, nil
}

func (m *ToolManager) SelectList(page, pageSize int) ([]*model.Tool, error) {
	pageOffset := app.GetPageOffset(page, pageSize)
	if pageOffset < 0 && pageSize < 0 {
		pageOffset = 0
		pageSize = 5
	}
	var tools []*model.Tool
	if err := m.conn.Offset(pageOffset).Limit(pageSize).Find(&tools).Error; err != nil {
		return nil, err
	}
	return tools, nil
}

func NewToolManager(table string, conn *gorm.DB) ITool {
	return &ToolManager{table: table, conn: conn}
}

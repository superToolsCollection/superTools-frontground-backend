package tools

import (
	"github.com/gin-gonic/gin"
	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/service"
	"superTools-frontground-backend/pkg/app"
	"superTools-frontground-backend/pkg/errcode"
)

/**
* @Author: super
* @Date: 2020-12-25 21:16
* @Description:
**/

type ToolController struct {
	ToolService service.IToolService
}

func NewToolController(toolService service.IToolService) ToolController {
	return ToolController{ToolService: toolService}
}

// @Summary 根据id获取工具
// @Tags tool
// @Produce json
// @Param id body string true "工具id"
// @Success 200 {object} service.Tool "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tools/getTool [get]
func (t ToolController) GetToolByKey(c *gin.Context) {
	param := service.GetToolByKeyRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	tool, err := t.ToolService.GetToolByKey(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetToolByKey err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetToolByKeyFail)
		return
	}
	response.ToResponse(tool)
	return
}

// @Summary 根据名称获取工具
// @Tags tool
// @Produce json
// @Param name body string true "工具名称"
// @Success 200 {object} service.Tool "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tools/getToolByName [get]
func (t ToolController) GetToolByName(c *gin.Context) {
	param := service.GetToolByNameRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	tool, err := t.ToolService.GetToolByName(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetToolByName err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetToolByNameFail)
		return
	}
	response.ToResponse(tool)
	return
}

// @Summary 获取工具列表
// @Tags tool
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} service.Tool "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tools/toolList [get]
func (t ToolController) GetToolList(c *gin.Context) {
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	response := app.NewResponse(c)
	tools, err := t.ToolService.GetToolList(&pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetToolList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetToolListFail)
		return
	}
	response.ToResponse(tools)
	return
}

// @Summary 获取工具列表
// @Tags tool
// @Produce json
// @Success 200 {object} service.Tool "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tools/allTools [get]
func (t ToolController) GetAllTools(c *gin.Context) {
	response := app.NewResponse(c)
	tools, err := t.ToolService.GetAllTools()
	if err != nil {
		global.Logger.Errorf(c, "svc.GetToolList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetToolListFail)
		return
	}
	response.ToResponse(tools)
	return
}

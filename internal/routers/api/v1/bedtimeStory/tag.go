package bedtimeStory

import (
	"github.com/gin-gonic/gin"

	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/service"
	"superTools-frontground-backend/pkg/app"
	"superTools-frontground-backend/pkg/errcode"
)

/**
* @Author: super
* @Date: 2020-09-16 07:50
* @Description: tag对应的restful api
**/

type TagController struct {
	TagService service.ITagService
}

func NewTagController(tagService service.ITagService) TagController {
	return TagController{TagService: tagService}
}

// @Summary 根据ID获取标签
// @tags 睡前故事
// @Produce  json
// @Param id query string false "标签id"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag/{id} [get]
func (t TagController) GetTag(c *gin.Context) {
	param := service.SelectTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	tag, err := t.TagService.GetTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	response.ToResponse(tag)
	return
}

// @Summary 根据ID获取多个标签
// @tags 睡前故事
// @Produce  json
// @Param id query string false "标签id"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t TagController) GetTags(c *gin.Context) {
	param := service.SelectTagsRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	tag, err := t.TagService.GetTags(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	response.ToResponse(tag)
	return
}

package bedtimeStory

import (
	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/service"
	"superTools-frontground-backend/pkg/app"
	"superTools-frontground-backend/pkg/convert"
	"superTools-frontground-backend/pkg/errcode"

	"github.com/gin-gonic/gin"
)

/**
* @Author: super
* @Date: 2020-09-16 07:51
* @Description: story对应的restful api
**/

type Story struct {
}

func NewStory() Story {
	return Story{}
}

// @Summary 获取单个故事
// @tags 睡前故事
// @Produce json
// @Param id path int true "故事ID"
// @Success 200 {object} model.BedtimeStory "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/bedtime/stories/{id} [get]
func (t Story) Get(c *gin.Context) {
	param := service.StoryRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	story, err := svc.GetStory(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetStory err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetStoryFail)
		return
	}

	response.ToResponse(story)
	return
}

// @Summary 仅获得单个故事的内容和作者
// @tags 睡前故事
// @Produce json
// @Param id path int true "故事ID"
// @Success 200 {object} model.BedtimeStory "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/bedtime/stories_only/{id} [get]
func (t Story) GetOnly(c *gin.Context) {
	param := service.StoryRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	story, err := svc.GetStoryOnly(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetStoryOnly err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetStoryFail)
		return
	}

	response.ToResponse(story)
	return
}

// @Summary 获取多个故事
// @tags 睡前故事
// @Produce json
// @Param tag_id query int false "标签ID"
// @Param state query int false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.BedtimeStorySwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/bedtime/stories [get]
func (t Story) List(c *gin.Context) {
	param := service.StoryListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	stories, totalRows, err := svc.GetStoryList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetStoryList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetStoriesFail)
		return
	}

	response.ToResponseList(stories, totalRows)
	return
}

// @Summary 创建故事
// @tags 睡前故事
// @Produce json
// @Param tag_id body string true "标签ID"
// @Param content body string true "故事内容"
// @Param author body string true "作者"
// @Param created_by body int true "创建者"
// @Param state body int false "状态"
// @Success 200 {object} model.BedtimeStory "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/bedtime/stories [post]
func (t Story) Create(c *gin.Context) {
	param := service.CreateStoryRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateStory(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateStory err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateStoryFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新故事
// @tags 睡前故事
// @Produce json
// @Param tag_id body string false "标签ID"
// @Param content body string false "故事内容"
// @Param modified_by body string true "修改者"
// @Success 200 {object} model.BedtimeStory "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/bedtime/stories/{id} [put]
func (t Story) Update(c *gin.Context) {
	param := service.UpdateStoryRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateStory(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateStory err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateStoryFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 删除故事
// @tags 睡前故事
// @Produce  json
// @Param id path int true "故事ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/bedtime/stories/{id} [delete]
func (t Story) Delete(c *gin.Context) {
	param := service.DeleteStoryRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteStory(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteStory err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteStoryFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

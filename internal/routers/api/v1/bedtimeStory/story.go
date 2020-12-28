package bedtimeStory

import (
	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/service"
	"superTools-frontground-backend/pkg/app"
	"superTools-frontground-backend/pkg/errcode"

	"github.com/gin-gonic/gin"
)

/**
* @Author: super
* @Date: 2020-09-16 07:51
* @Description: story对应的restful api
**/

type StoryController struct {
	StoryService service.IStoryService
}

func NewStoryController(storyService service.IStoryService) StoryController {
	return StoryController{StoryService: storyService}
}

// @Summary 随机获取单个故事
// @tags 睡前故事
// @Produce json
// @Success 200 {object} model.BedtimeStory "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/bedtime/story [get]
func (t StoryController) Get(c *gin.Context) {
	response := app.NewResponse(c)
	story, err := t.StoryService.GetStory()
	if err != nil {
		global.Logger.Errorf(c, "svc.GetStory err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetStoryFail)
		return
	}

	response.ToResponse(story)
	return
}

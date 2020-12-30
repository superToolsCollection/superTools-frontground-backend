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
* @Date: 2020-12-30 13:43
* @Description:
**/
type ForbesController struct {
	ForbesService service.IForbesService
}

func NewForbesController(forbesService service.IForbesService) ForbesController {
	return ForbesController{ForbesService: forbesService}
}

// @Summary 获取福布斯排行榜
// @tags tool
// @Produce  json
// @success 200 {object} service.Forbes "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/forbes/all [get]
func (f ForbesController) GetForbes(c *gin.Context) {
	response := app.NewResponse(c)
	forbes, err := f.ForbesService.GetForbes()
	if err != nil {
		global.Logger.Errorf(c, "get all forbes err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetForbesFail)
		return
	}
	response.ToResponse(forbes)
}

// @Summary 获取福布斯排行榜
// @tags tool
// @Produce  json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @success 200 {object} service.Forbes "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/forbes/list [get]
func (f ForbesController) GetForbesList(c *gin.Context) {
	response := app.NewResponse(c)
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	forbes, err := f.ForbesService.GetForbesList(&pager)
	if err != nil {
		global.Logger.Errorf(c, "get forbes list err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetForbesListFail)
		return
	}
	response.ToResponse(forbes)
}

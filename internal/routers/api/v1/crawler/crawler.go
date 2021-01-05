package crawler

import (
	"github.com/gin-gonic/gin"
	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/service"
	"superTools-frontground-backend/pkg/app"
	"superTools-frontground-backend/pkg/errcode"
)

type CrawlController struct {
	ForbesService service.IForbesService
}

func NewForbesController(forbesService service.IForbesService) ForbesController {
	return ForbesController{ForbesService: forbesService}
}

/**
* @Author: super
* @Date: 2020-12-31 14:51
* @Description:
**/

// @Summary 获取福布斯排行榜
// @tags tool
// @Produce  json
// @success 200 {object} service.Forbes "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/crawler/data [get]
func CrawlData(c *gin.Context) {
	response := app.NewResponse(c)
	forbes, err := f.ForbesService.GetForbes()
	if err != nil {
		global.Logger.Errorf(c, "get all forbes err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetForbesFail)
		return
	}
	response.ToResponse(forbes)
}

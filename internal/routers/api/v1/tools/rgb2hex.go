package tools

import (
	"github.com/gin-gonic/gin"
	"superTools-frontground-backend/global"
	"superTools-frontground-backend/pkg/app"
	"superTools-frontground-backend/pkg/errcode"
	"superTools-frontground-backend/pkg/util"
)

/**
* @Author: super
* @Date: 2020-11-27 16:35
* @Description:
**/

// @Summary 根据传递的字符串生成二维码字节的base64编码
// @tags tool
// @Produce  json
// @Param str query string true "要转换为hex的rgb字符串"
// @success 200 {string} string "#xxxxxx"
// @Router /api/v1/rgb2hex [get]
func RgbToHex(c *gin.Context) {
	response := app.Response{c}
	str := c.Query("str")
	result, err := util.RgbToHex(str)
	if err != nil {
		global.Logger.Errorf(c, "rgb to hex err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails())
		return
	}
	response.ToResponse(result)
}

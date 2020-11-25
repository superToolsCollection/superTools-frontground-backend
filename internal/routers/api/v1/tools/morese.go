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
* @Date: 2020-08-24 10:12
* @Description:
**/

// @Summary 根据传递的字符串生成摩尔斯密码
// @tags tool
// @Produce  json
// @Param str query string true "要生成摩尔斯密码的字符串"
// @success 200 {string} string "{"data":{}}"
// @Router /api/v1/morse [get]
func GetMorse(c *gin.Context) {
	//appG := app.Gin{C: c}
	response := app.NewResponse(c)
	str := c.Query("str")

	s, err := util.GenerateMorse(str)

	if err != nil {
		global.Logger.Errorf(c, "generate morse err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails())
		//appG.Response(http.StatusOK, errcode.INVALID_PARAMS, "")
		return
	}

	response.ToResponse(s)
	//appG.Response(http.StatusOK, errcode.Success, s)
}

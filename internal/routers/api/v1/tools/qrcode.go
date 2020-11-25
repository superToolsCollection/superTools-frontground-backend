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
* @Date: 2020-08-21 21:50
* @Description:
**/

// @Summary 根据传递的字符串生成二维码字节的base64编码
// @tags tool
// @Produce  json
// @Param str query string true "要生成对应二维码的地址"
// @success 200 {string} string "{"data":{}}"
// @Router /api/v1/qrcode [get]
func GetQRcode(c *gin.Context) {
	//appG := app.Gin{C: c}
	response := app.Response{c}
	str := c.Query("str")

	bytes, err := util.GenerateQRCodeByte(str)

	if err != nil {
		global.Logger.Errorf(c, "generate qr-code err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails())
		//appG.Response(http.StatusOK, errcode.INVALID_PARAMS, "")
		return
	}

	encode, err := util.EncodeBase64(string(bytes))
	response.ToResponse(encode)

	//appG.Response(http.StatusOK, errcode.SUCCESS, encode)
}

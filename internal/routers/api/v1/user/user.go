package user

import (
	"github.com/gin-gonic/gin"

	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/service"
	"superTools-frontground-backend/pkg/app"
	"superTools-frontground-backend/pkg/errcode"
)

/**
* @Author: super
* @Date: 2020-11-24 18:37
* @Description:
**/

type UserController struct {
	UserService service.IUserService
}

func NewUserController(userService service.IUserService) UserController {
	return UserController{UserService: userService}
}

// @Summary 用户登录
// @Tags user
// @Produce json
// @Param user_name body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} service.User "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/user/signin [post]
func (u UserController) SignIn(c *gin.Context) {
	//todo:cookie与session设置
	param := service.UserSignInRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, err := u.UserService.SignIn(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.SignIn err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserSignInFail)
		return
	}
	response.ToResponse(user)
	return
}

// @Summary 更新用户信息
// @Tags user
// @Produce json
// @Param id body string true "用户id"
// @Param user_name body string true "用户名"
// @Param password body string true "密码"
// @Param nick_name body string true "昵称"
// @Success 200 {object} string "success"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/orders [put]
func (u UserController) Update(c *gin.Context) {
	param := service.UserUpdateInfoRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	err := u.UserService.UpdateInfo(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateInfo err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserUpdateFail)
		return
	}
	response.ToResponse("success")
	return
}

// @Summary 用户注册
// @Tags user
// @Produce json
// @Param user_name body string true "用户名"
// @Param password body string true "密码"
// @Param nick_name body string true "昵称"
// @Success 200 {object} string "userID"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/orders [post]
func (u UserController) Register(c *gin.Context) {
	param := service.UserRegisterRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	userId, err := u.UserService.Register(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.Register err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserRegisterFail)
		return
	}
	response.ToResponse(userId)
	return
}

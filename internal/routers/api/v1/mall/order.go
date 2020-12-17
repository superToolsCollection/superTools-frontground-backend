package mall

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"superTools-frontground-backend/internal/model"
	"superTools-frontground-backend/pkg/util"

	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/service"
	"superTools-frontground-backend/pkg/app"
	"superTools-frontground-backend/pkg/convert"
	"superTools-frontground-backend/pkg/errcode"
)

/**
* @Author: super
* @Date: 2020-11-21 15:57
* @Description:
**/

type OrderController struct {
	OrderService service.IOrderService
}

func NewOrderController(orderService service.IOrderService) OrderController {
	return OrderController{OrderService: orderService}
}

// @Summary 获取订单列表
// @Tags mall
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} service.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/orders [get]
func (o OrderController) GetOrderList(c *gin.Context) {
	response := app.NewResponse(c)
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	orders, err := o.OrderService.GetOrderList(&pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetOrderList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOrderListFail)
		return
	}
	response.ToResponse(orders)
	return
}

// @Summary 获取所有订单
// @Tags mall
// @Produce json
// @Success 200 {object} service.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/all_orders [get]
func (o OrderController) GetAllOrder(c *gin.Context) {
	response := app.NewResponse(c)
	orders, err := o.OrderService.GetAllOrder()
	if err != nil {
		global.Logger.Errorf(c, "svc.GetAllOrder err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetAllOrderFail)
		return
	}
	response.ToResponse(orders)
	return
}

// @Summary 获取用户所有订单
// @Tags mall
// @Produce json
// @Param user_id query int false "用户id"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} service.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/all_orders_user [get]
func (o OrderController) GetOrderListByUserID(c *gin.Context) {
	param := service.GetOrderListByUserIDRequest{}
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	orders, err := o.OrderService.GetOrderListByUserID(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetOrderListByUserID err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOrderListByUserIDFail)
		return
	}
	response.ToResponse(orders)
	return
}

// @Summary 获取用户订单列表
// @Tags mall
// @Produce json
// @Param user_id query int false "用户id"
// @Success 200 {object} service.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/orders_user [get]
func (o OrderController) GetOrderByUserID(c *gin.Context) {
	param := service.GetOrderByUserIDRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	orders, err := o.OrderService.GetOrderByUserID(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetOrderByUserID err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOrderByUserIDFail)
		return
	}
	response.ToResponse(orders)
	return
}

// @Summary 获取单个订单
// @Tags mall
// @Produce json
// @Param id path int true "订单ID"
// @Success 200 {object} service.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/orders/{id} [get]
func (o OrderController) GetOrder(c *gin.Context) {
	param := service.OrderRequest{ID: convert.StrTo(c.Param("id")).MustInt64()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	order, err := o.OrderService.GetOrderByID(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetOrderByID err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOrderFail)
		return
	}
	response.ToResponse(order)
	return
}

// @Summary 新增订单
// @Tags mall
// @Produce json
// @Param id body int true "订单id"
// @Param user_id body int true "用户id"
// @Param product_id body int true "商品id"
// @Param state body int true "订单状态"
// @Success 200 {object} int "1"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/orders [post]
func (o OrderController) Insert(c *gin.Context) {
	param := service.InsertOrderRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	userInfoJson, err := c.Cookie("loginUserJson")
	//todo:用户禁用cookie的解决方案
	if err != nil {
		global.Logger.Errorf(c, "svc.InsertOrder err: %v", err)
		response.ToErrorResponse(errcode.ErrorInsertOrderFail)
		return
	}
	decodeStruct, err := util.DecodeToStruct(userInfoJson)
	userInfo := decodeStruct.(service.LoginUser)

	message := model.NewMessage(param.ProductID, userInfo.ID)
	messageJson, err := json.Marshal(message)
	if err != nil {
		//todo:修改报错信息
		global.Logger.Errorf(c, "svc.InsertOrder err: %v", err)
		response.ToErrorResponse(errcode.ErrorInsertOrderFail)
		return
	}
	//todo:将消息发送到消息队列
	fmt.Println(messageJson)

	return
}

// @Summary 更新订单
// @Tags mall
// @Produce json
// @Param id body int true "订单id"
// @Param user_id body int true "用户id"
// @Param product_id body int true "商品id"
// @Param state body int true "订单状态"
// @Success 200 {object} string "success"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/orders [put]
func (o OrderController) Update(c *gin.Context) {
	param := service.UpdateOrderRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	err := o.OrderService.UpdateOrder(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateOrder err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateOrderFail)
		return
	}
	response.ToResponse("success")
	return
}

// @Summary 删除订单
// @Tags mall
// @Produce json
// @Param id path int true "订单ID"
// @Success 200 {string} string "success"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/orders/{id} [delete]
func (o OrderController) Delete(c *gin.Context) {
	param := service.DeleteOrderRequest{ID: convert.StrTo(c.Param("id")).MustInt64()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	result := o.OrderService.DeleteOrderByID(&param)
	if result != true {
		global.Logger.Errorf(c, "svc.DeleteOrderByID err: %v", result)
		response.ToErrorResponse(errcode.ErrorDeleteOrderFail)
		return
	}
	response.ToResponse("success")
	return
}

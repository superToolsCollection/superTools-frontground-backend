package errcode

/**
* @Author: super
* @Date: 2020-11-20 10:56
* @Description:
**/

var (
	ErrorGetProductFail     = NewError(20040001, "获取单个商品失败")
	ErrorGetAllProductFail  = NewError(20040002, "获取全部商品失败")
	ErrorGetProductListFail = NewError(20040003, "获取多个商品失败")
	ErrorInsertProductFail  = NewError(20040004, "插入商品失败")
	ErrorUpdateProductFail  = NewError(20040005, "更新商品失败")
	ErrorDeleteProductFail  = NewError(20040006, "删除商品失败")

	ErrorGetOrderFail             = NewError(20050001, "获取单个订单失败")
	ErrorGetAllOrderFail          = NewError(20050002, "获取全部订单失败")
	ErrorGetOrderListFail         = NewError(20050003, "获取多个订单失败")
	ErrorInsertOrderFail          = NewError(20050004, "插入订单失败")
	ErrorUpdateOrderFail          = NewError(20050005, "更新订单失败")
	ErrorDeleteOrderFail          = NewError(20050006, "删除订单失败")
	ErrorGetOrderByUserIDFail     = NewError(20050007, "删除用户订单失败")
	ErrorGetOrderListByUserIDFail = NewError(20050008, "删除用户订单列表失败")
)

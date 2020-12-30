package errcode

/**
* @Author: super
* @Date: 2020-11-25 22:44
* @Description:
**/

var (
	ErrorGetToolByKeyFail  = NewError(20070001, "根据ID获取工具信息失败")
	ErrorGetToolByNameFail = NewError(20070002, "根据名称获取工具失败")
	ErrorGetToolListFail   = NewError(20070003, "获取工具列表失败")

	ErrorGetForbesFail     = NewError(20070004, "获取福布斯排行榜错误")
	ErrorGetForbesListFail = NewError(20070005, "获取福布斯排行榜列表错误")
)

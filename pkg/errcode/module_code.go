package errcode

/**
* @Author: super
* @Date: 2020-09-22 09:49
* @Description: 统一错误代码
**/

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010004, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")

	ErrorGetStoryFail    = NewError(20020001, "获取单个故事失败")
	ErrorGetStoriesFail  = NewError(20020002, "获取多个故事失败")
	ErrorCreateStoryFail = NewError(20020003, "创建故事失败")
	ErrorUpdateStoryFail = NewError(20020004, "更新故事失败")
	ErrorDeleteStoryFail = NewError(20020005, "删除故事失败")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")
)

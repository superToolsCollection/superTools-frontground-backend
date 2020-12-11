package mall

import (
	"github.com/gin-gonic/gin"

	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/service"
	"superTools-frontground-backend/pkg/app"
	"superTools-frontground-backend/pkg/convert"
	"superTools-frontground-backend/pkg/errcode"
)

/**
* @Author: super
* @Date: 2020-11-18 16:34
* @Description:
**/

type ProductController struct {
	ProductService service.IProductService
}

func NewProductController(productService service.IProductService) ProductController {
	return ProductController{productService}
}

// @Summary 获取商品列表
// @Tags mall
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} service.Product "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/products [get]
func (p ProductController) GetProductList(c *gin.Context) {
	response := app.NewResponse(c)
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	products, err := p.ProductService.GetProductList(&pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetProductList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetProductListFail)
		return
	}
	response.ToResponse(products)
	return
}

// @Summary 获取所有商品
// @Tags mall
// @Produce json
// @Success 200 {object} service.Product "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/all_product [get]
func (p ProductController) GetAllProduct(c *gin.Context) {
	response := app.NewResponse(c)
	products, err := p.ProductService.GetAllProduct()
	if err != nil {
		global.Logger.Errorf(c, "svc.GetAllProduct err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetAllProductFail)
		return
	}
	response.ToResponse(products)
	return
}

// @Summary 获取单个商品
// @Tags mall
// @Produce json
// @Param id path int true "商品ID"
// @Success 200 {object} service.Product "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/products/{id} [get]
func (p ProductController) GetProduct(c *gin.Context) {
	param := service.ProductRequest{ID: convert.StrTo(c.Param("id")).MustInt64()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	product, err := p.ProductService.GetProductByID(param.ID)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetProduct err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetProductFail)
		return
	}
	response.ToResponse(product)
	return
}

// @Summary 新增商品
// @Tags mall
// @Produce json
// @Param id body int true "商品ID"
// @Param product_name body string true "商品名称"
// @Param product_num body int true "商品数量"
// @Param product_image body string true "商品图像"
// @Param product_url body string true "商品链接"
// @Success 200 {object} int "1"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/products [post]
func (p ProductController) Insert(c *gin.Context) {
	param := service.InsertProductRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	product, err := p.ProductService.InsertProduct(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.InsertProduct err: %v", err)
		response.ToErrorResponse(errcode.ErrorInsertProductFail)
		return
	}
	response.ToResponse(product)
	return
}

// @Summary 更新商品
// @Tags mall
// @Produce json
// @Param id body int true "商品ID"
// @Param product_name body string true "商品名称"
// @Param product_num body int true "商品数量"
// @Param product_image body string true "商品图像"
// @Param product_url body string true "商品链接"
// @Success 200 {object} string "success"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/products [put]
func (p ProductController) Update(c *gin.Context) {
	param := service.UpdateProductRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	err := p.ProductService.UpdateProduct(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateProduct err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateProductFail)
		return
	}
	response.ToResponse("success")
	return
}

// @Summary 删除商品
// @Tags mall
// @Produce json
// @Param id path int true "商品ID"
// @Success 200 {string} string "success"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/products/{id} [delete]
func (p ProductController) Delete(c *gin.Context) {
	param := service.DeleteProductRequest{ID: convert.StrTo(c.Param("id")).MustInt64()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	result := p.ProductService.DeleteProductByID(&param)
	if result != true {
		global.Logger.Errorf(c, "svc.DeleteProduct err: %v", result)
		response.ToErrorResponse(errcode.ErrorDeleteProductFail)
		return
	}
	response.ToResponse("success")
	return
}

// @Summary 获取静态文件
// @Tags mall
// @Produce json
// @Param id path int true "商品ID"
// @Success 200 {string} string "success"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/mall/GetGenerateHtml [get]
func (p ProductController) GetGenerateHtml(c *gin.Context) {
	param := service.ProductRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	result, err := p.ProductService.GetProductByID(param.ID)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetGenerateHtml err: %v", result)
		response.ToErrorResponse(errcode.ErrorGetGenerateHtmlFail)
		return
	}
	response.ToResponse("success")
	return
}
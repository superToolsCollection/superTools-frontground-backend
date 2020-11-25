package service

import (
	"fmt"

	"superTools-frontground-backend/internal/dao"
	"superTools-frontground-backend/pkg/app"
)

/**
* @Author: super
* @Date: 2020-11-18 15:40
* @Description:
**/
type ProductRequest struct {
	ID int64 `form:"id" binding:"required,gte=1"`
}

type InsertProductRequest struct {
	ID           int64  `form:"id" binding:"required,gte=1"`
	ProductName  string `form:"product_name" binding:"required,min=2,max=4294967295"`
	ProductNum   int64  `form:"product_num" binding:"required,gte=1"`
	ProductImage string `form:"product_image" binding:"required,min=2,max=1000"`
	ProductUrl   string `form:"product_url" binding:"required,min=2,max=4294967295"`
}

type UpdateProductRequest struct {
	ID           int64  `form:"id" binding:"required,gte=1"`
	ProductName  string `form:"product_name" binding:"required,min=2,max=4294967295"`
	ProductNum   int64  `form:"product_num" binding:"required,gte=1"`
	ProductImage string `form:"product_image" binding:"required,min=2,max=1000"`
	ProductUrl   string `form:"product_url" binding:"required,min=2,max=4294967295"`
}

type DeleteProductRequest struct {
	ID int64 `form:"id" binding:"required,gte=1"`
}

type Product struct {
	ID           int64  `json:"id"`
	ProductName  string `json:"ProductName"`
	ProductNum   int64  `json:"ProductNum"`
	ProductImage string `json:"ProductImage"`
	ProductUrl   string `json:"ProductUrl"`
}

func (p *Product) String() string {
	return fmt.Sprintf("id: %d", p.ID)
}

type IProductService interface {
	GetProductByID(int64) (*Product, error)
	GetAllProduct() ([]*Product, error)
	GetProductList(pager *app.Pager) ([]*Product, error)
	DeleteProductByID(param *DeleteProductRequest) bool
	InsertProduct(param *InsertProductRequest) (int64, error)
	UpdateProduct(param *UpdateProductRequest) error
}

type ProductService struct {
	productDao dao.IProduct
}

//初始化函数
func NewProductService(productDao dao.IProduct) IProductService {
	return &ProductService{productDao}
}

func (p *ProductService) GetProductByID(productID int64) (*Product, error) {
	product, err := p.productDao.SelectByKey(productID)
	if err != nil {
		return nil, err
	}
	return &Product{
		ID:           product.ID,
		ProductName:  product.ProductName,
		ProductUrl:   product.ProductUrl,
		ProductImage: product.ProductImage,
		ProductNum:   product.ProductNum,
	}, nil
}

func (p *ProductService) GetAllProduct() ([]*Product, error) {
	products, err := p.productDao.SelectAll()
	if err != nil {
		return nil, err
	}
	var productList []*Product
	for _, product := range products {
		productList = append(productList, &Product{
			ID:           product.ID,
			ProductName:  product.ProductName,
			ProductUrl:   product.ProductUrl,
			ProductImage: product.ProductImage,
			ProductNum:   product.ProductNum,
		})
	}
	return productList, nil
}

func (p *ProductService) GetProductList(pager *app.Pager) ([]*Product, error) {
	products, err := p.productDao.SelectList(pager.Page, pager.PageSize)
	if err != nil {
		return nil, err
	}
	var productList []*Product
	for _, product := range products {
		productList = append(productList, &Product{
			ID:           product.ID,
			ProductName:  product.ProductName,
			ProductUrl:   product.ProductUrl,
			ProductImage: product.ProductImage,
			ProductNum:   product.ProductNum,
		})
	}
	return productList, nil
}

func (p *ProductService) DeleteProductByID(param *DeleteProductRequest) bool {
	return p.productDao.Delete(param.ID)
}

func (p *ProductService) InsertProduct(param *InsertProductRequest) (int64, error) {
	product := &dao.Product{
		ID:           param.ID,
		ProductName:  param.ProductName,
		ProductNum:   param.ProductNum,
		ProductImage: param.ProductImage,
		ProductUrl:   param.ProductUrl,
	}
	return p.productDao.Insert(product)
}

func (p *ProductService) UpdateProduct(param *UpdateProductRequest) error {
	product := &dao.Product{
		ID:           param.ID,
		ProductName:  param.ProductName,
		ProductNum:   param.ProductNum,
		ProductImage: param.ProductImage,
		ProductUrl:   param.ProductUrl,
	}
	err := p.productDao.Update(product)
	if err != nil {
		return err
	}
	return nil
}

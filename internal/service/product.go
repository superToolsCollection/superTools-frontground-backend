package service

import (
	"fmt"
	"os"
	"path/filepath"
	"superTools-frontground-backend/internal/model"
	"text/template"

	"superTools-frontground-backend/internal/dao"
	"superTools-frontground-backend/pkg/app"
)

/**
* @Author: super
* @Date: 2020-11-18 15:40
* @Description:
**/
var (
	htmlOutPath  = "/Users/super/develop/superTools-frontground-backend/web/htmlProductShow/" //生成的html文件保存位置
	templatePath = "/Users/super/develop/superTools-frontground-backend/web/views/template/"  //静态文件模板目录
)

type ProductRequest struct {
	ID string `form:"id" binding:"required,min=2,max=100"`
}

type InsertProductRequest struct {
	ID           string `form:"id" binding:"required,min=2,max=100"`
	ProductName  string `form:"product_name" binding:"required,min=2,max=4294967295"`
	ProductNum   int64  `form:"product_num" binding:"required,gte=1"`
	ProductImage string `form:"product_image" binding:"required,min=2,max=1000"`
	ProductUrl   string `form:"product_url" binding:"required,min=2,max=4294967295"`
}

type UpdateProductRequest struct {
	ID           string `form:"id" binding:"required,min=2,max=100"`
	ProductName  string `form:"product_name" binding:"required,min=2,max=4294967295"`
	ProductNum   int64  `form:"product_num" binding:"required,gte=1"`
	ProductImage string `form:"product_image" binding:"required,min=2,max=1000"`
	ProductUrl   string `form:"product_url" binding:"required,min=2,max=4294967295"`
}

type DeleteProductRequest struct {
	ID string `form:"id" binding:"required,min=2,max=100"`
}

type Product struct {
	ID           string `json:"id"`
	ProductName  string `json:"ProductName"`
	ProductNum   int64  `json:"ProductNum"`
	ProductImage string `json:"ProductImage"`
	ProductUrl   string `json:"ProductUrl"`
}

func (p *Product) String() string {
	return fmt.Sprintf("id: %d", p.ID)
}

type IProductService interface {
	GetProductByID(string) (*Product, error)
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

func (p *ProductService) GetProductByID(productID string) (*Product, error) {
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

//用于实现页面静态化
func (p *ProductService) GetGenerateHtml(param *ProductRequest) error {
	//1.获取模版
	contenstTmp, err := template.ParseFiles(filepath.Join(templatePath, "product.html"))
	if err != nil {
		return err
	}
	//2.获取html生成路径
	fileName := filepath.Join(htmlOutPath, "htmlProduct.html")

	//3.获取模版渲染数据
	product, err := p.productDao.SelectByKey(param.ID)
	if err != nil {
		return err
	}
	//4.生成静态文件
	return generateStaticHtml(contenstTmp, fileName, product)

}

//生成html静态文件
func generateStaticHtml(template *template.Template, fileName string, product *model.Product) error {
	//1.判断静态文件是否存在
	if exist(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			return err
		}
	}
	//2.生成静态文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	err = template.Execute(file, &product)
	if err != nil {
		return err
	}
	return nil
}

//判断文件是否存在
func exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

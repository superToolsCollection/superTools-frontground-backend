package dao

import (
	"errors"

	"superTools-frontground-backend/internal/model"
	"superTools-frontground-backend/pkg/app"

	"github.com/jinzhu/gorm"
)

/**
* @Author: super
* @Date: 2020-11-18 14:41
* @Description: 显示商品相关信息
**/

type Product struct {
	ID           string `json:"id" sql:"ID"`
	ProductName  string `json:"ProductName"`
	ProductNum   int64  `json:"ProductNum"`
	ProductImage string `json:"ProductImage"`
	ProductUrl   string `json:"ProductUrl"`
}

type IProduct interface {
	Insert(product *Product) (int64, error)
	Delete(id string) bool
	Update(product *Product) error
	SelectByKey(string) (*model.Product, error)
	SelectAll() ([]*model.Product, error)
	SelectList(page, pageSize int) ([]*model.Product, error)
}

type ProductManager struct {
	table string
	conn  *gorm.DB
}

func NewProductManager(table string, conn *gorm.DB) IProduct {
	return &ProductManager{table: table, conn: conn}
}

func (p *ProductManager) Insert(product *Product) (int64, error) {
	productCreate := &model.Product{
		ID:           product.ID,
		ProductName:  product.ProductName,
		ProductNum:   product.ProductNum,
		ProductImage: product.ProductImage,
		ProductUrl:   product.ProductUrl,
	}
	result := p.conn.Create(productCreate)
	if result.RowsAffected == int64(0) {
		return 0, errors.New("insert error")
	}
	return result.RowsAffected, nil
}

func (p *ProductManager) Update(product *Product) error {
	productUpdate := &model.Product{
		ID:           product.ID,
		ProductName:  product.ProductName,
		ProductNum:   product.ProductNum,
		ProductImage: product.ProductImage,
		ProductUrl:   product.ProductUrl,
	}
	result := p.conn.Model(productUpdate).Where("id = ?", product.ID).Updates(product)

	if result.RowsAffected == int64(0) {
		return errors.New("update error")
	}
	return nil
}

func (p *ProductManager) Delete(id string) bool {
	result := p.conn.Where("id = ?", id).Delete(model.Product{})
	if result.RowsAffected == int64(0) {
		return false
	}
	return true
}

func (p *ProductManager) SelectByKey(id string) (*model.Product, error) {
	product := &model.Product{}
	result := p.conn.Where("id = ?", id).Find(product)
	if result.RecordNotFound() {
		return nil, errors.New("wrong id")
	}
	return product, nil
}

func (p *ProductManager) SelectAll() ([]*model.Product, error) {
	var products []*model.Product
	if err := p.conn.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductManager) SelectList(page, pageSize int) ([]*model.Product, error) {
	pageOffset := app.GetPageOffset(page, pageSize)
	if pageOffset < 0 && pageSize < 0 {
		pageOffset = 0
		pageSize = 5
	}
	fields := []string{"id", "product_name", "product_num", "product_image", "product_url"}
	rows, err := p.conn.Offset(pageOffset).Limit(pageSize).Select(fields).Table(p.table).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		product := &model.Product{}
		if err := rows.Scan(&product.ID,
			&product.ProductName,
			&product.ProductNum,
			&product.ProductImage,
			&product.ProductUrl); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

package service

import (
	"github.com/jinzhu/gorm"

	"strings"
	"testing"

	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/dao"
	"superTools-frontground-backend/pkg/app"
	"superTools-frontground-backend/pkg/db"
	"superTools-frontground-backend/pkg/setting"
)

/**
* @Author: super
* @Date: 2020-11-18 16:25
* @Description:
**/

func GetConn() (*gorm.DB, error) {
	newSetting, err := setting.NewSetting(strings.Split("/Users/super/develop/superTools-backend/configs", ",")...)
	if err != nil {
		return nil, err
	}
	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return nil, err
	}
	conn, err := db.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func TestProductService_GetProductByID(t *testing.T) {
	conn, err := GetConn()
	if err != nil {
		t.Error(err)
	}
	productManager := dao.NewProductManager("products", conn)
	productService := NewProductService(productManager)
	product, err := productService.GetProductByID(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(product)
}

func TestProductService_GetAllProduct(t *testing.T) {
	conn, err := GetConn()
	if err != nil {
		t.Error(err)
	}
	productManager := dao.NewProductManager("products", conn)
	productService := NewProductService(productManager)
	products, err := productService.GetAllProduct()
	if err != nil {
		t.Error(err)
	}
	t.Log(products)
}

func TestProductService_GetProductList(t *testing.T) {
	conn, err := GetConn()
	if err != nil {
		t.Error(err)
	}
	productManager := dao.NewProductManager("products", conn)
	productService := NewProductService(productManager)
	products, err := productService.GetProductList(&app.Pager{
		Page:     2,
		PageSize: 2,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(products)
}

func TestProductService_InsertProduct(t *testing.T) {
	conn, err := GetConn()
	if err != nil {
		t.Error(err)
	}
	productManager := dao.NewProductManager("products", conn)
	productService := NewProductService(productManager)

	product := &InsertProductRequest{
		ID:          6,
		ProductName: "test",
		ProductNum:  12,
		ProductUrl:  "444",
	}

	products, err := productService.InsertProduct(product)
	if err != nil {
		t.Error(err)
	}
	t.Log(products)
}

func TestProductService_UpdateProduct(t *testing.T) {
	conn, err := GetConn()
	if err != nil {
		t.Error(err)
	}
	productManager := dao.NewProductManager("products", conn)
	productService := NewProductService(productManager)

	product := &UpdateProductRequest{
		ID:          6,
		ProductName: "test",
		ProductNum:  12,
		ProductUrl:  "888",
	}

	err = productService.UpdateProduct(product)
	if err != nil {
		t.Error(err)
	}
}

func TestProductService_DeleteProductByID(t *testing.T) {
	conn, err := GetConn()
	if err != nil {
		t.Error(err)
	}
	productManager := dao.NewProductManager("products", conn)
	productService := NewProductService(productManager)
	id := &DeleteProductRequest{
		ID: 6,
	}
	isDelete := productService.DeleteProductByID(id)
	t.Log(isDelete)
}

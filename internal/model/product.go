package model

/**
* @Author: super
* @Date: 2020-11-18 14:19
* @Description: 商品类
* @Group: 秒杀
**/

type Product struct {
	ID           int64  `gorm:"column:id;primary_key" json:"id"           sql:"ID"           consite:"ID"`
	ProductName  string `gorm:"column:product_name"   json:"ProductName"  sql:"productName"  consite:"ProductName"`
	ProductNum   int64  `gorm:"column:product_num"    json:"ProductNum"   sql:"productNum"   consite:"ProductNum"`
	ProductImage string `gorm:"column:product_image"  json:"ProductImage" sql:"productImage" consite:"ProductImage"`
	ProductUrl   string `gorm:"column:product_url"    json:"ProductUrl"   sql:"productUrl"   consite:"ProductUrl"`
}

func (p Product) TableName() string {
	return "products"
}

package model

/**
* @Author: super
* @Date: 2020-11-21 10:42
* @Description: 订单类
* @Group: 秒杀
**/

const (
	OrderWait    = iota //0为待支付状态
	OrderSuccess        //1为下单成功
	OrderFailed         //2为下单失败
)

type Order struct {
	ID        int64  `gorm:"column:id;primary_key" json:"id"         sql:"id"`
	UserID    string `gorm:"column:user_id"        json:"user_id"    sql:"user_id"`
	ProductID int64  `gorm:"column:product_id"     json:"product_id" sql:"product_id"`
	State     int    `gorm:"column:state"          json:"state"      sql:"state"`
}

func (o Order) TableName() string {
	return "orders"
}
